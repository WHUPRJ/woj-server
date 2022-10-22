package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func (h *handler) Build(_ context.Context, t *asynq.Task) error {
	// TODO: configure timeout with context

	var p model.ProblemBuildPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	h.log.Info("build", zap.Any("payload", p))

	config, status := h.runnerService.NewProblem(p.ProblemVersionID, p.ProblemFile)

	for i := range config.Languages {
		config.Languages[i].Type = ""
		config.Languages[i].Script = ""
		config.Languages[i].Cmp = ""
	}

	b, _ := json.Marshal(config)
	h.taskService.ProblemUpdate(status, p.ProblemVersionID, string(b))

	return nil
}
