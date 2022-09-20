package user

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"gorm.io/gorm"
)

func (s *service) Profile(id uint) (*model.User, e.Status) {
	user := new(model.User)

	err := s.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, e.UserNotFound
	}
	if err != nil {
		return user, e.DatabaseError
	}

	return user, e.Success
}
