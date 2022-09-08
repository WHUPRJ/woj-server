package user

import (
	"github.com/WHUPRJ/woj-server/internal/repo/postgresql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type CreateData struct {
	Username string
	Nickname string
	Password []byte
}

func (s *service) Create(data *CreateData) (id uint, err error) {
	model := &postgresql.User{
		UserName:  data.Username,
		Password:  data.Password,
		NickName:  data.Nickname,
		IsEnabled: true,
	}

	if err = s.db.Get().Create(model).Error; err != nil {
		s.log.Debug("create user error", zap.Error(err), zap.Any("data", data))
		return 0, errors.Wrap(err, "create error")
	}
	return model.ID, nil
}
