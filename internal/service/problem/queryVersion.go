package problem

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *service) QueryVersion(pvid uint, shouldEnable bool) (*model.ProblemVersion, e.Status) {
	problemVersion := new(model.ProblemVersion)

	err := s.db.First(&problemVersion, pvid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ProblemVersionNotFound
	}
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("pvid", pvid))
		return nil, e.DatabaseError
	}

	if shouldEnable && !problemVersion.IsEnabled {
		return nil, e.ProblemVersionNotAvailable
	}
	return problemVersion, e.Success
}
