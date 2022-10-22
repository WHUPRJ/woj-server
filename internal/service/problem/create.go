package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

type CreateData struct {
	Title      string
	Statement  string
	ProviderID uint
	IsEnabled  bool
}

func (s *service) Create(data *CreateData) (*model.Problem, e.Status) {
	problem := &model.Problem{
		Title:      data.Title,
		Statement:  data.Statement,
		ProviderID: data.ProviderID,
		IsEnabled:  data.IsEnabled,
	}

	err := s.db.Create(problem).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("problem", problem))
		return nil, e.DatabaseError
	}

	return problem, e.Success
}
