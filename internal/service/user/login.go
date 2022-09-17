package user

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(data *model.User) (*model.User, e.Err) {
	user := &model.User{UserName: data.UserName}

	err := s.db.Get().Where(user).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, e.UserNotFound
	}
	if err != nil {
		return user, e.DatabaseError
	}

	err = bcrypt.CompareHashAndPassword(user.Password, data.Password)
	if err != nil {
		return user, e.UserWrongPassword
	}

	return user, e.Success
}
