package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type CreateData struct {
	Username string
	Nickname string
	Password string
}

func (s *service) Create(data *CreateData) (uint, e.Err) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Debug("bcrypt error", zap.Error(err), zap.String("password", data.Password))
		return 0, e.InternalError
	}

	user := &model.User{
		UserName:  data.Username,
		Password:  hashed,
		NickName:  data.Nickname,
		IsEnabled: true,
	}

	if err := s.db.Get().Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return 0, e.UserDuplicated
		}
		s.log.Debug("create user error", zap.Error(err), zap.Any("data", data))
		return 0, e.DatabaseError
	}
	return user.ID, e.Success
}
