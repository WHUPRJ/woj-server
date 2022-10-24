package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type CreateData struct {
	UserName string
	Password string
	NickName string
}

func (s *service) Create(data *CreateData) (*model.User, e.Status) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Warn("BcryptError", zap.Error(err), zap.String("password", data.Password))
		return nil, e.InternalError
	}

	user := &model.User{
		UserName:  data.UserName,
		Password:  hashed,
		NickName:  data.NickName,
		Role:      model.RoleGeneral,
		IsEnabled: true,
	}

	err = s.db.Create(user).Error
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return nil, e.UserDuplicated
	}
	if err != nil {
		s.log.Warn("DatabaseError", zap.Error(err), zap.Any("user", user))
		return nil, e.DatabaseError
	}
	return user, e.Success
}
