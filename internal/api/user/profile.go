package user

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type profileRequest struct {
	UID uint `form:"uid"`
}

// Profile
// @Summary     profile
// @Description fetch user profile
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       uid formData int false "user id"
// @Response    200 {object} e.Response "user info"
// @Security    Authentication
// @Router      /v1/user/profile [post]
func (h *handler) Profile(c *gin.Context) {
	// TODO: create a new struct for profile (user info & solve info)

	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	uid := claim.(*global.Claim).UID
	role := claim.(*global.Claim).Role
	req := new(profileRequest)
	if err := c.ShouldBind(req); err == nil {
		if req.UID != 0 && req.UID != uid {
			if role >= model.RoleGeneral {
				uid = req.UID
			} else {
				e.Pong(c, e.UserUnauthorized, nil)
				return
			}
		}
	}

	user, status := h.userService.Profile(uid)
	e.Pong(c, status, user)
}
