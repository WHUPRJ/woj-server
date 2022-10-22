package task

import (
	"encoding/json"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/runner"
	"go.uber.org/zap"
)

func (s *service) SubmitJudge(pvid uint, storageKey string, submission model.Submission) (string, e.Status) {
	payload, err := json.Marshal(
		model.SubmitJudgePayload{
			ProblemVersionId: pvid,
			StorageKey:       storageKey,
			Submission:       submission,
		})
	if err != nil {
		s.log.Warn("json marshal error", zap.Error(err), zap.Any("Submission", submission))
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeSubmitJudge, payload, model.QueueRunner)

	return info.ID, status
}

func (s *service) SubmitUpdate(status e.Status, sid uint, point int32, ctx runner.JudgeStatus) (string, e.Status) {
	ctxText, err := json.Marshal(ctx)
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("ctx", ctx))
		return "", e.InternalError
	}

	payload, err := json.Marshal(model.SubmitUpdatePayload{
		Status:  status,
		Sid:     sid,
		Point:   point,
		Context: string(ctxText),
	})
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("Status", status),
			zap.Int32("Point", point),
			zap.Any("Context", ctx))
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeSubmitUpdate, payload, model.QueueServer)

	return info.ID, status
}
