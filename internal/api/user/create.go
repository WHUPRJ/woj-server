package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
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
// @Response    200 {object} e.Response "random string"
// @Security    Authentication
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

	id, err := h.userService.Create(createData)
	e.Pong(c, err, id)
}
