package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `form:"username" binding:"required"`
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
	userData := &model.User{
		UserName: req.Username,
		Password: []byte(req.Password),
	}
	user, status := h.userService.Login(userData)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	// sign and return token
	version, status := h.userService.IncrVersion(user.ID)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}
	claim := &global.Claim{
		UID:     user.ID,
		Role:    user.Role,
		Version: version,
	}
	token, status := h.jwtService.SignClaim(claim)
	e.Pong(c, status, gin.H{
		"token":    token,
		"nickname": user.NickName,
	})
}
