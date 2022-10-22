package task

import (
	"encoding/json"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) ProblemBuild(pvId uint, file string) (string, e.Status) {
	payload, err := json.Marshal(model.ProblemBuildPayload{
		ProblemVersionID: pvId,
		ProblemFile:      file,
	})
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("ProblemVersionID", pvId),
			zap.String("ProblemFile", file))
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeProblemBuild, payload, model.QueueRunner)

	return info.ID, status
}

func (s *service) ProblemUpdate(status e.Status, pvId uint, ctx string) (string, e.Status) {
	payload, err := json.Marshal(model.ProblemUpdatePayload{
		Status:           status,
		ProblemVersionID: pvId,
		Context:          ctx,
	})
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("Status", status),
			zap.Any("ProblemVersionID", pvId),
			zap.Any("Context", ctx))
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeProblemUpdate, payload, model.QueueServer)

	return info.ID, status
}
