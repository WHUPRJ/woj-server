package user

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(data *CreateData) (id uint, err error)
}

type service struct {
	log *zap.Logger
	db  global.Repo
}

func NewUserService(g *global.Global) Service {
	return &service{
		log: g.Log,
		db:  g.Db,
	}
}
