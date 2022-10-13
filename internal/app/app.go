package app

import (
	"context"
	"fmt"
	"github.com/WHUPRJ/woj-server/global"
	"github.com/WHUPRJ/woj-server/internal/router"
	"github.com/WHUPRJ/woj-server/internal/service/jwt"
	"github.com/WHUPRJ/woj-server/repo/postgresql"
	"github.com/WHUPRJ/woj-server/repo/redis"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(g *global.Global) error {
	// Setup Database
	g.Db = new(postgresql.Repo)
	g.Db.Setup(g)

	// Setup Redis
	g.Redis = new(redis.Repo)
	g.Redis.Setup(g)

	// Setup JWT
	g.Jwt = jwt.NewJwtService(g)

	// Prepare Router
	handler := router.InitRouters(g)

	// Create Server
	addr := fmt.Sprintf("%s:%d", g.Conf.WebServer.Address, g.Conf.WebServer.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Run Server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.Log.Fatal("ListenAndServe Failed", zap.Error(err))
		}
	}()

	// Handle SIGINT and SIGTERM.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	g.Log.Info("Shutting down server ...")

	// Graceful Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		g.Log.Warn("Server Shutdown Failed", zap.Error(err))
	}
	err = g.Db.Close()
	if err != nil {
		g.Log.Warn("Database Close Failed", zap.Error(err))
	}

	return err
}
