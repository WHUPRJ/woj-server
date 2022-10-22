package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func (s *service) Query(pid uint, uid uint, offset int, limit int) ([]*model.Submission, e.Status) {
	submissions := make([]*model.Submission, 0)

	submission := &model.Submission{
		ProblemID: pid,
		UserID:    uid,
	}

	err := s.db.Preload(clause.Associations).
		Where(submission).
		Limit(limit).
		Offset(offset).
		Find(&submissions).Error

	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return nil, e.ProblemNotFound
	//}

	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("pid", pid), zap.Any("uid", uid))
		return nil, e.DatabaseError
	}
	return submissions, e.Success
}
