package jwt

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	ParseToken(tokenText string) (*Claim, e.Err)
	SignClaim(claim *Claim) (string, e.Err)
	// TODO: Validate(claim *Claim) bool

	Handler() gin.HandlerFunc
}

type service struct {
	log        *zap.Logger
	SigningKey []byte
	ExpireHour int
}

func NewJwtService(g *global.Global) Service {
	return &service{
		log:        g.Log,
		SigningKey: []byte(g.Conf.WebServer.JwtSigningKey),
		ExpireHour: g.Conf.WebServer.JwtExpireHour,
	}
}
