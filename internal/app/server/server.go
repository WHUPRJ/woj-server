package server

import (
	"context"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/api/consumer"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/repo/postgresql"
	"github.com/WHUPRJ/woj-server/internal/repo/redis"
	"github.com/WHUPRJ/woj-server/internal/router"
	"github.com/WHUPRJ/woj-server/internal/service/jwt"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/WHUPRJ/woj-server/pkg/zapasynq"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func RunServer(g *global.Global) error {
	// Setup Database
	g.Db = new(postgresql.Repo)
	g.Db.Setup(g)

	// Setup Redis
	g.Redis = new(redis.Repo)
	g.Redis.Setup(g)

	// Setup JWT
	g.Jwt = jwt.NewJwtService(g)

	// Prepare Router
	routers := router.InitRouters(g)

	// Create Server
	addr := fmt.Sprintf("%s:%d", g.Conf.WebServer.Address, g.Conf.WebServer.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: routers,
	}

	// Run Server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.Log.Fatal("ListenAndServe Failed", zap.Error(err))
		}
	}()

	// Create Queue
	queueMux := asynq.NewServeMux()
	{
		handler := consumer.NewConsumer(g)
		queueMux.HandleFunc(model.TypeProblemUpdate, handler.ProblemUpdate)
		queueMux.HandleFunc(model.TypeSubmitUpdate, handler.SubmitUpdate)
	}
	queueSrv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     g.Conf.Redis.Address,
			Password: g.Conf.Redis.Password,
			DB:       g.Conf.Redis.QueueDb,
		},
		asynq.Config{
			Concurrency: utils.If(runtime.NumCPU() > 1, runtime.NumCPU()-1, 1).(int),
			Logger:      zapasynq.New(g.Log),
			Queues:      map[string]int{model.QueueServer: 1},
		},
	)

	// Run Queue
	if err := queueSrv.Start(queueMux); err != nil {
		g.Log.Fatal("queueSrv.Start Failed", zap.Error(err))
	}

	// Handle SIGINT and SIGTERM.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	g.Log.Info("Shutting down server ...")

	// Graceful Shutdown Server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		g.Log.Warn("Server Shutdown Failed", zap.Error(err))
	}

	// Graceful Shutdown Queue
	queueSrv.Shutdown()

	// Graceful Shutdown Database
	err = g.Db.Close()
	if err != nil {
		g.Log.Warn("Database Close Failed", zap.Error(err))
	}

	return err
}
