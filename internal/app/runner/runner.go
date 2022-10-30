package runner

import (
	"github.com/WHUPRJ/woj-server/internal/api/runner"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/WHUPRJ/woj-server/pkg/zapasynq"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"runtime"
)

func RunRunner(g *global.Global) error {
	hnd, err := runner.NewRunner(g)
	if err != nil {
		return err
	}

	mux := asynq.NewServeMux()
	mux.HandleFunc(model.TypeProblemBuild, hnd.Build)
	mux.HandleFunc(model.TypeSubmitJudge, hnd.Judge)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     g.Conf.Redis.Address,
			Password: g.Conf.Redis.Password,
			DB:       g.Conf.Redis.QueueDb,
		},
		asynq.Config{
			Concurrency: utils.If(runtime.NumCPU() > 1, runtime.NumCPU()-1, 1),
			Logger:      zapasynq.New(g.Log),
			Queues:      map[string]int{model.QueueRunner: 1},
		},
	)

	if err := srv.Run(mux); err != nil {
		g.Log.Warn("could not run server", zap.Error(err))
		return err
	}

	return nil
}
