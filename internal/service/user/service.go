package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(data *CreateData) (*model.User, e.Status)
	Login(data *LoginData) (*model.User, e.Status)
	IncrVersion(uid uint) (int64, e.Status)
	Profile(uid uint) (*model.User, e.Status)
}

type service struct {
	log   *zap.Logger
	db    *gorm.DB
	redis *redis.Client
}

func NewService(g *global.Global) Service {
	return &service{
		log:   g.Log,
		db:    g.Db.Get().(*gorm.DB),
		redis: g.Redis.Get().(*redis.Client),
	}
}
