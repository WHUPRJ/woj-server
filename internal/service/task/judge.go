package task

import (
	"encoding/json"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) NewJudge(submission model.Submission) (string, e.Status) {
	payload, err := json.Marshal(model.SubmitJudge{Submission: submission})
	if err != nil {
		s.log.Warn("json marshal error", zap.Error(err), zap.Any("payload", submission))
		return "", e.InternalError
	}

	info, status := s.submit(model.TypeSubmitJudge, payload)

	return info.ID, status
}
