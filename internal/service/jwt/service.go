package jwt

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var _ global.JwtService = (*service)(nil)

type service struct {
	log        *zap.Logger
	redis      *redis.Client
	SigningKey []byte
	ExpireHour int
}

func NewJwtService(g *global.Global) global.JwtService {
	return &service{
		log:        g.Log,
		redis:      g.Redis.Get().(*redis.Client),
		SigningKey: []byte(g.Conf.WebServer.JwtSigningKey),
		ExpireHour: g.Conf.WebServer.JwtExpireHour,
	}
}
