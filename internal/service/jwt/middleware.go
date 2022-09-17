package jwt

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
	"strings"
)

func (s *service) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(strings.ToLower(tokenHeader), "bearer ") {
			e.Pong(c, e.TokenEmpty, nil)
			c.Abort()
			return
		}
		token := tokenHeader[7:]

		claim, err := s.ParseToken(token)
		if err != e.Success {
			e.Pong(c, err, nil)
			c.Abort()
			return
		}

		// TODO: validate claim version
		// if !s.Validate(claim) {
		// 	e.Pong(c, e.TokenRevoked, nil)
		// 	c.Abort()
		// }

		c.Set("claim", claim)
		c.Next()
	}
}
