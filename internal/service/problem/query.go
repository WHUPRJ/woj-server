package problem

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *service) Query(problemId uint) (*model.Problem, e.Status) {
	problem := new(model.Problem)

	err := s.db.Preload(clause.Associations).First(&problem, problemId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ProblemNotFound
	}
	if err != nil {
		return nil, e.DatabaseError
	}

	return problem, e.Success
}