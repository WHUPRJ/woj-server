package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(data *CreateData) (*model.User, e.Status)
	Login(data *model.User) (*model.User, e.Status)
	IncrVersion(id uint) (int64, e.Status)
	Profile(id uint) (*model.User, e.Status)
}

type service struct {
	log   *zap.Logger
	db    *gorm.DB
	redis *redis.Client
}

func NewUserService(g *global.Global) Service {
	return &service{
		log:   g.Log,
		db:    g.Db.Get().(*gorm.DB),
		redis: g.Redis.Get().(*redis.Client),
	}
}
