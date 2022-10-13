package user

import (
	"github.com/WHUPRJ/woj-server/global"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
)

// Logout
// @Summary     logout
// @Description logout
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Response    200 {object} e.Response "nil"
// @Security    Authentication
// @Router      /v1/user/logout [post]
func (h *handler) Logout(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	_, status := h.userService.IncrVersion(claim.(*global.Claim).UID)
	e.Pong(c, status, nil)
}
