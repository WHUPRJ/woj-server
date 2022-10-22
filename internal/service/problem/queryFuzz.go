package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func (s *service) QueryFuzz(search string, associations bool, shouldEnable bool) ([]*model.Problem, e.Status) {
	problems := make([]*model.Problem, 0)

	query := s.db
	if associations {
		query = query.Preload(clause.Associations)
	}
	if shouldEnable {
		query = query.Where("is_enabled = true")
	}
	query = query.
		Where(s.db.Where("title LIKE ?", "%"+search+"%").
			Or("statement LIKE ?", "%"+search+"%"))
	err := query.Find(&problems).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("search", search))
		return nil, e.DatabaseError
	}

	return problems, e.Success
}
