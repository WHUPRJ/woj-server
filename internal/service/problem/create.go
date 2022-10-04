package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/model"
	"go.uber.org/zap"
)

func (s *service) Create(uid uint, problem *model.Problem) (*model.Problem, e.Status) {
	problem.ProviderID = uid
	problem.IsEnabled = true

	if err := s.db.Create(problem).Error; err != nil {
		s.log.Debug("create problem error", zap.Error(err), zap.Any("problem", problem))
		return nil, e.DatabaseError
	}

	return problem, e.Success
}
