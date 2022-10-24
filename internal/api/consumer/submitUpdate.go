package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/status"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func (h *handler) SubmitUpdate(_ context.Context, t *asynq.Task) error {
	p := new(model.SubmitUpdatePayload)
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	if p.Status != e.Success {
		h.log.Warn("RunnerError", zap.Any("payload", p))
		return nil
	}

	createData := &status.CreateData{
		SubmissionID:     p.SubmissionID,
		ProblemVersionID: p.ProblemVersionID,
		Context:          p.Context,
		Point:            p.Point,
	}
	_, eStatus := h.statusService.Create(createData)

	if eStatus != e.Success {
		return fmt.Errorf(eStatus.String())
	}
	return nil
}
