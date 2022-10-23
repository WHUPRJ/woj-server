package consumer

import (
	"context"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/problem"
	"github.com/WHUPRJ/woj-server/internal/service/status"
	"github.com/WHUPRJ/woj-server/internal/service/task"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	ProblemUpdate(_ context.Context, t *asynq.Task) error
	SubmitUpdate(_ context.Context, t *asynq.Task) error
}

type handler struct {
	log            *zap.Logger
	problemService problem.Service
	statusService  status.Service
	taskService    task.Service
}

func NewConsumer(g *global.Global) Handler {
	hnd := &handler{
		log:            g.Log,
		problemService: problem.NewService(g),
		statusService:  status.NewService(g),
		taskService:    task.NewService(g),
	}

	return hnd
}
