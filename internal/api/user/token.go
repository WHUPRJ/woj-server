package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) tokenNext(c *gin.Context, user *model.User) {
	claim := &global.Claim{
		UID:      user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
	}
	token, err := h.jwtService.SignClaim(claim)
	e.Pong(c, err, token)
}
