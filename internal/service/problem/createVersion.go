package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/jackc/pgtype"
	"go.uber.org/zap"
)

type CreateVersionData struct {
	ProblemID  uint
	StorageKey string
}

func (s *service) CreateVersion(data *CreateVersionData) (*model.ProblemVersion, e.Status) {
	problemVersion := &model.ProblemVersion{
		ProblemID:  data.ProblemID,
		Context:    pgtype.JSON{Status: pgtype.Null},
		StorageKey: data.StorageKey,
	}

	err := s.db.Create(problemVersion).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("problemVersion", problemVersion))
		return nil, e.DatabaseError
	}

	return problemVersion, e.Success
}
