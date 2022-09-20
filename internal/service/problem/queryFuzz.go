package problem

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"gorm.io/gorm"
)

func (s *service) QueryFuzz(search string) ([]*model.Problem, e.Status) {
	var problems []*model.Problem

	err := s.db.
		Where("title LIKE ?", "%"+search+"%").
		Or("content LIKE ?", "%"+search+"%").
		Find(&problems).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ProblemNotFound
	}
	if err != nil {
		return nil, e.DatabaseError
	}

	return problems, e.Success
}
