package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type createRequest struct {
	Username string `form:"username" binding:"required"`
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Create godoc
// @Summary     create a new user
// @Description create a new user
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       username formData string true "username"
// @Param       nickname formData string true "nickname"
// @Param       password formData string true "password"
// @Response    200 {object} e.Response "random string"
// @Router      /v1/user [post]
func (h *handler) Create(c *gin.Context) {
	req := new(createRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.log.Debug("bcrypt error", zap.Error(err), zap.String("password", req.Password))
		e.Pong(c, e.InternalError, err.Error())
		return
	}

	createData := new(user.CreateData)
	createData.Nickname = req.Nickname
	createData.Username = req.Username
	createData.Password = hashed

	id, err := h.service.Create(createData)
	if err != nil {
		e.Pong(c, e.DatabaseError, err.Error())
		return
	}

	e.Pong(c, e.Success, id)
}
