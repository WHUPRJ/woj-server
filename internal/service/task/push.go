package task

import (
	"encoding/json"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) PushProblem(id uint, file string) (string, e.Status) {
	payload, err := json.Marshal(model.ProblemPushPayload{
		ProblemID:   id,
		ProblemFile: file,
	})
	if err != nil {
		s.log.Warn("json marshal error", zap.Error(err), zap.Any("id", id), zap.String("file", file))
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeSubmitJudge, payload)

	return info.ID, status
}
