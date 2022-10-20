package user

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(data *model.User) (*model.User, e.Status) {
	user := &model.User{UserName: data.UserName}

	err := s.db.Where(user).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.UserNotFound
	}
	if err != nil {
		return nil, e.DatabaseError
	}

	if !user.IsEnabled {
		return nil, e.UserDisabled
	}

	err = bcrypt.CompareHashAndPassword(user.Password, data.Password)
	if err != nil {
		return nil, e.UserWrongPassword
	}

	return user, e.Success
}
