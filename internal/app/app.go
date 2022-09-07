package app

import (
	"context"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(g *global.Global) error {
	routersInit := router.InitRouters(g)

	addr := fmt.Sprintf("%s:%d", g.Conf.WebServer.Address, g.Conf.WebServer.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: routersInit,
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		g.Log.Fatal("Server Shutdown Failed", zap.Error(err))
	}

	return err
}
