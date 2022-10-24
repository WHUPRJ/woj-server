package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

type CreateData struct {
	ProblemID uint
	UserID    uint
	Language  string
	Code      string
}

func (s *service) Create(data *CreateData) (*model.Submission, e.Status) {
	submission := &model.Submission{
		ProblemID: data.ProblemID,
		UserID:    data.UserID,
		Language:  data.Language,
		Code:      data.Code,
	}

	err := s.db.Create(submission).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("submission", submission))
		return nil, e.DatabaseError
	}

	return submission, e.Success
}
