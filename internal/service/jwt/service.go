package jwt

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"go.uber.org/zap"
)

var _ global.JwtService = (*service)(nil)

type service struct {
	log        *zap.Logger
	SigningKey []byte
	ExpireHour int
}

func NewJwtService(g *global.Global) global.JwtService {
	return &service{
		log:        g.Log,
		SigningKey: []byte(g.Conf.WebServer.JwtSigningKey),
		ExpireHour: g.Conf.WebServer.JwtExpireHour,
	}
}
