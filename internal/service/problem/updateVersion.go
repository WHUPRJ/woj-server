package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
)

func (s *service) UpdateVersion(pvid uint, values interface{}) e.Status {
	err := s.db.Model(&model.ProblemVersion{}).Where("id = ?", pvid).Updates(values).Error
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("pvid", pvid), zap.Any("values", values))
		return e.DatabaseError
	}

	return e.Success
}
