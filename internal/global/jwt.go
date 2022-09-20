package global

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	UID     uint       `json:"id"`
	Role    model.Role `json:"role"`
	Version int64      `json:"version"`
	jwt.RegisteredClaims
}

type JwtService interface {
	ParseToken(tokenText string) (*Claim, e.Status)
	SignClaim(claim *Claim) (string, e.Status)
	Validate(claim *Claim) bool

	Handler() gin.HandlerFunc
}
