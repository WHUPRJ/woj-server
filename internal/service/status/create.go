package status

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/jackc/pgtype"
	"go.uber.org/zap"
)

type CreateData struct {
	SubmissionID     uint
	ProblemVersionID uint
	Context          pgtype.JSON
	Point            int32
}

func (s service) Create(data *model.Status) (*model.Status, e.Status) {
	status := &model.Status{
		SubmissionID:     data.SubmissionID,
		ProblemVersionID: data.ProblemVersionID,
		Context:          data.Context,
		Point:            data.Point,
		IsEnabled:        true,
	}

	err := s.db.Create(status).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("status", status))
		return nil, e.DatabaseError
	}

	return status, e.Success
}
