package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) UpdateVersion(problemVersion *model.ProblemVersion) (*model.ProblemVersion, e.Status) {
	err := s.db.Save(problemVersion).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("problemVersion", problemVersion))
		return nil, e.DatabaseError
	}

	return problemVersion, e.Success
}
