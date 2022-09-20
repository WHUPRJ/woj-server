package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/user"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Username string `form:"username" binding:"required"`
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Create
// @Summary     create a new user
// @Description create a new user
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       username formData string true "username"
// @Param       nickname formData string true "nickname"
// @Param       password formData string true "password"
// @Response    200 {object} e.Response "jwt token"
// @Router      /v1/user/create [post]
func (h *handler) Create(c *gin.Context) {
	req := new(createRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	createData := &user.CreateData{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	}

	u, status := h.userService.Create(createData)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	version, status := h.userService.IncrVersion(u.ID)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}
	claim := &global.Claim{
		UID:     u.ID,
		Role:    u.Role,
		Version: version,
	}
	token, status := h.jwtService.SignClaim(claim)
	e.Pong(c, status, token)
}
