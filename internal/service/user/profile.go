package user

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *service) Profile(uid uint) (*model.User, e.Status) {
	user := new(model.User)

	err := s.db.First(&user, uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.UserNotFound
	}
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("uid", uid))
		return nil, e.DatabaseError
	}

	return user, e.Success
}
