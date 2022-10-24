package task

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/runner"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	ProblemBuild(data *model.ProblemBuildPayload) (string, e.Status)
	ProblemUpdate(data *model.ProblemUpdatePayload) (string, e.Status)
	SubmitJudge(data *model.SubmitJudgePayload) (string, e.Status)
	SubmitUpdate(data *model.SubmitUpdatePayload, ctx runner.JudgeStatus) (string, e.Status)

	GetTaskInfo(string, string) (*asynq.TaskInfo, e.Status)
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
		DB:       g.Conf.Redis.QueueDb,
	}
	return &service{
		log:       g.Log,
		queue:     asynq.NewClient(redisOpt),
		inspector: asynq.NewInspector(redisOpt),
	}
}
