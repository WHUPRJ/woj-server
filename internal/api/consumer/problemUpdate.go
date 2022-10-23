package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgtype"
	"go.uber.org/zap"
)

func (h *handler) ProblemUpdate(_ context.Context, t *asynq.Task) error {
	p := new(model.ProblemUpdatePayload)
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	if p.Status != e.Success {
		h.log.Warn("RunnerError", zap.Any("payload", p))
		return nil
	}

	status := h.problemService.UpdateVersion(
		p.ProblemVersionID,
		map[string]interface{}{
			"Context": pgtype.JSON{
				Bytes:  []byte(p.Context),
				Status: pgtype.Present,
			},
			"IsEnabled": true,
		},
	)

	if status != e.Success {
		return fmt.Errorf(status.String())
	}
	return nil
}
