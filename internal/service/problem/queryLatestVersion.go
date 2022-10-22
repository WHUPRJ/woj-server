package problem

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *service) QueryLatestVersion(pid uint) (*model.ProblemVersion, e.Status) {
	problemVersion := &model.ProblemVersion{
		ProblemID: pid,
		IsEnabled: true,
	}

	err := s.db.
		Where(problemVersion).
		Last(&problemVersion).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ProblemVersionNotFound
	}
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("problemVersion", problemVersion))
		return nil, e.DatabaseError
	}

	return problemVersion, e.Success
}
