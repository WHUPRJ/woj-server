package problem

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *service) Query(pid uint, associations bool, shouldEnable bool) (*model.Problem, e.Status) {
	problem := new(model.Problem)

	query := s.db
	if associations {
		query = query.Preload(clause.Associations)
	}
	err := query.First(&problem, pid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ProblemNotFound
	}
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("pid", pid))
		return nil, e.DatabaseError
	}

	if shouldEnable && !problem.IsEnabled {
		return nil, e.ProblemNotAvailable
	}
	return problem, e.Success
}
