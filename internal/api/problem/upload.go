package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// Upload
// @Summary     get upload url
// @Description get upload url
// @Produce     json
// @Response    200 {object} e.Response "upload url and key"
// @Security    Authentication
// @Router      /v1/problem/upload [post]
func (h *handler) Upload(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	role := claim.(*global.Claim).Role
	if role < model.RoleAdmin {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	key := utils.RandomString(16)
	url, status := h.storageService.Upload(key, time.Second*60*60)

	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	e.Pong(c, e.Success, gin.H{
		"key": key,
		"url": url,
	})
}
