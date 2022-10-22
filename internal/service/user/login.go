package user

import (
	"errors"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginData struct {
	UserName string
	Password string
}

func (s *service) Login(data *LoginData) (*model.User, e.Status) {
	user := &model.User{UserName: data.UserName}

	err := s.db.Where(user).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.UserNotFound
	}
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("user", user))
		return nil, e.DatabaseError
	}

	if !user.IsEnabled {
		return nil, e.UserDisabled
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password))
	if err != nil {
		return nil, e.UserWrongPassword
	}

	return user, e.Success
}
