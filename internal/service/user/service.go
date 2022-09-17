package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(data *CreateData) (uint, e.Err)
	Login(data *model.User) (*model.User, e.Err)
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
