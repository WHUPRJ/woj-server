package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) Update(problem *model.Problem) (*model.Problem, e.Status) {
	err := s.db.Save(problem).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("problem", problem))
		return nil, e.DatabaseError
	}

	return problem, e.Success
}
