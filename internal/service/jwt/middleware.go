package jwt

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/gin-gonic/gin"
	"strings"
)

func (s *service) Handler(forced bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, status := func() (*global.Claim, e.Status) {
			const tokenPrefix = "bearer "
			tokenHeader := c.GetHeader("Authorization")
			if tokenHeader == "" || !strings.HasPrefix(strings.ToLower(tokenHeader), tokenPrefix) {
				return nil, e.TokenEmpty
			}

			token := tokenHeader[len(tokenPrefix):]
			claim, status := s.ParseToken(token)
			if status != e.Success {
				return nil, status
			}

			if !s.Validate(claim) {
				return nil, e.TokenRevoked
			}
			return claim, e.Success
		}()

		if status == e.Success {
			c.Set("claim", claim)
		}
		if forced && status != e.Success {
			e.Pong(c, status, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
