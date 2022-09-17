package global

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	UID      uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Version  int    `json:"version"`
	jwt.RegisteredClaims
}

type JwtService interface {
	ParseToken(tokenText string) (*Claim, e.Err)
	SignClaim(claim *Claim) (string, e.Err)
	// TODO: Validate(claim *Claim) bool

	Handler() gin.HandlerFunc
}
