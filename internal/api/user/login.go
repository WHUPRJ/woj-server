package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
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
// @Response    200 {object} e.Response "random string"
// @Router      /v1/user/login [post]
func (h *handler) Login(c *gin.Context) {
	req := new(loginRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	// check password
	userData := &model.User{
		UserName: req.Username,
		Password: []byte(req.Password),
	}
	user, err := h.userService.Login(userData)
	if err != e.Success {
		e.Pong(c, err, nil)
		return
	}

	// sign and return token
	h.tokenNext(c, user)
}
