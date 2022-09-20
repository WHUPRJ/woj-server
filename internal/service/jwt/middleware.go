package jwt

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
	"strings"
)

func (s *service) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		const tokenPrefix = "bearer "
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(strings.ToLower(tokenHeader), tokenPrefix) {
			e.Pong(c, e.TokenEmpty, nil)
			c.Abort()
			return
		}
		token := tokenHeader[len(tokenPrefix):]

		claim, status := s.ParseToken(token)
		if status != e.Success {
			e.Pong(c, status, nil)
			c.Abort()
			return
		}

		if !s.Validate(claim) {
			e.Pong(c, e.TokenRevoked, nil)
			c.Abort()
			return
		}

		c.Set("claim", claim)
		c.Next()
	}
}
