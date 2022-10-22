package runner

import (
	"context"
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/runner"
	"github.com/WHUPRJ/woj-server/internal/service/task"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Build(_ context.Context, t *asynq.Task) error
	Judge(_ context.Context, t *asynq.Task) error
}

type handler struct {
	log           *zap.Logger
	runnerService runner.Service
	taskService   task.Service
}

func NewRunner(g *global.Global) (Handler, error) {
	hnd := &handler{
		log:           g.Log,
		runnerService: runner.NewService(g),
		taskService:   task.NewService(g),
	}

	status := hnd.runnerService.EnsureDeps(false)
	if status != e.Success {
		g.Log.Error("failed to ensure runner dependencies", zap.String("status", status.String()))
		return nil, errors.New("failed to ensure dependencies")
	}

	return hnd, nil
}
