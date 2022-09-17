package user

import (
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type CreateData struct {
	Username string
	Nickname string
	Password string
}

func (s *service) Create(data *CreateData) (id uint, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Debug("bcrypt error", zap.Error(err), zap.String("password", data.Password))
		return 0, errors.Wrap(err, "bcrypt error")
	}

	user := &model.User{
		UserName:  data.Username,
		Password:  hashed,
		NickName:  data.Nickname,
		IsEnabled: true,
	}

	if err := s.db.Get().Create(user).Error; err != nil {
		s.log.Debug("create user error", zap.Error(err), zap.Any("data", data))
		return 0, errors.Wrap(err, "create error")
	}
	return user.ID, nil
}
