package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"time"
)

func (h *handler) Build(_ context.Context, t *asynq.Task) error {
	var p model.ProblemBuildPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	h.log.Info("build", zap.Any("payload", p))

	status, ctx := func() (e.Status, string) {
		url, status := h.storageService.Get(p.StorageKey, time.Second*60*5)
		if status != e.Success {
			return e.InternalError, "{}"
		}

		config, status := h.runnerService.NewProblem(p.ProblemVersionID, url, true)
		if status != e.Success {
			return e.InternalError, "{}"
		}

		for i := range config.Languages {
			config.Languages[i].Type = ""
			config.Languages[i].Script = ""
			config.Languages[i].Cmp = ""
		}

		b, _ := json.Marshal(config)
		return e.Success, string(b)
	}()

	h.taskService.ProblemUpdate(&model.ProblemUpdatePayload{
		Status:           status,
		ProblemVersionID: p.ProblemVersionID,
		Context:          ctx,
	})

	return nil
}
