package task

import (
	"encoding/json"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) ProblemBuild(data *model.ProblemBuildPayload) (string, e.Status) {
	payload, err := json.Marshal(data)
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("data", data),
		)
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeProblemBuild, payload, model.QueueRunner)

	return info.ID, status
}

func (s *service) ProblemUpdate(data *model.ProblemUpdatePayload) (string, e.Status) {
	payload, err := json.Marshal(data)
	if err != nil {
		s.log.Warn("json marshal error",
			zap.Error(err),
			zap.Any("data", data),
		)
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeProblemUpdate, payload, model.QueueServer)

	return info.ID, status
}
