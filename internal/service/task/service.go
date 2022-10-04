package task

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/model"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	NewJudge(submission model.Submission) e.Status
	submit(typename string, payload []byte) (*asynq.TaskInfo, e.Status)
}

type service struct {
	log       *zap.Logger
	queue     *asynq.Client
	inspector *asynq.Inspector
}

func NewService(g *global.Global) Service {
	redisOpt := asynq.RedisClientOpt{
		Addr:     g.Conf.Redis.Address,
		Password: g.Conf.Redis.Password,
		DB:       g.Conf.Redis.Db,
	}
	return &service{
		log:       g.Log,
		queue:     asynq.NewClient(redisOpt),
		inspector: asynq.NewInspector(redisOpt),
	}
}
