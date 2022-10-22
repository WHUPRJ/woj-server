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
	ProblemBuild(pvId uint, file string) (string, e.Status)
	ProblemUpdate(status e.Status, pvId uint, ctx string) (string, e.Status)
	SubmitJudge(pvid uint, storageKey string, submission model.Submission) (string, e.Status)
	SubmitUpdate(status e.Status, sid uint, point int32, ctx runner.JudgeStatus) (string, e.Status)

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
