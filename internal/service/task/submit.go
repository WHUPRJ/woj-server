package task

import (
	"encoding/json"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/runner"
	"go.uber.org/zap"
)

func (s *service) SubmitJudge(data *model.SubmitJudgePayload) (string, e.Status) {
	payload, err := json.Marshal(data)
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("data", data),
		)
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeSubmitJudge, payload, model.QueueRunner)

	return info.ID, status
}

func (s *service) SubmitUpdate(data *model.SubmitUpdatePayload, ctx runner.JudgeStatus) (string, e.Status) {
	ctxText, err := json.Marshal(ctx)
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("ctx", ctx))
		return "", e.InternalError
	}

	data.Context = string(ctxText)
	payload, err := json.Marshal(data)
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("data", data),
			zap.Any("Context", ctx),
		)
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeSubmitUpdate, payload, model.QueueServer)

	return info.ID, status
}
