package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/user"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Login
// @Summary     login
// @Description login and return token
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       username formData string true "username"
// @Param       password formData string true "password"
// @Response    200 {object} e.Response "jwt token and user nickname"
// @Router      /v1/user/login [post]
func (h *handler) Login(c *gin.Context) {
	req := new(loginRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, nil)
		return
	}

	// check password
	loginData := &user.LoginData{
		UserName: req.UserName,
		Password: req.Password,
	}
	u, status := h.userService.Login(loginData)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	// sign and return token
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
	e.Pong(c, status, gin.H{
		"token":    token,
		"nickname": u.NickName,
	})
}
