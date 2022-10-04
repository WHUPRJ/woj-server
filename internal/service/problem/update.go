package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func (s *service) Update(pid uint, problem *model.Problem) (*model.Problem, e.Status) {
	if err := s.db.Clauses(clause.Returning{}).Model(problem).
		Where("ID = (?)", pid).
		Select("Title", "Content", "TimeLimit", "MemoryLimit", "IsEnabled").
		Updates(problem).Error; err != nil {
		s.log.Debug("update problem error", zap.Error(err), zap.Any("problem", problem))
		return nil, e.DatabaseError
	}

	return problem, e.Success
}
