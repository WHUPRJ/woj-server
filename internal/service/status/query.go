package status

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s service) Query(sid uint, associations bool) (*model.Status, e.Status) {

	status := &model.Status{
		SubmissionID: sid,
		IsEnabled:    true,
	}

	query := s.db
	if associations {
		query = query.Preload(clause.Associations)
	}

	err := query.
		Where(status).
		Last(&status).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.StatusNotFound
	}

	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("status", status))
		return nil, e.DatabaseError
	}
	return status, e.Success
}

func (s service) QueryByVersion(pvid uint, offset int, limit int) ([]*model.Status, e.Status) {
	var statuses []*model.Status
	status := &model.Status{
		ProblemVersionID: pvid,
		IsEnabled:        true,
	}

	err := s.db.Preload(clause.Associations).
		Where(status).
		Limit(limit).
		Offset(offset).
		Find(&statuses).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("status", status))
		return nil, e.DatabaseError
	}
	return statuses, e.Success
}
