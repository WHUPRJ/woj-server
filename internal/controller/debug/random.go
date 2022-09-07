package debug

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// randomString godoc
// @Summary     random string
// @Description generate random string with length = 32
// @Tags        debug
// @Produce     json
// @Response    200 {object} e.Response "random string"
// @Router      /debug/random [get]
func (h *handler) randomString(c *gin.Context) {
	str := utils.RandomString(32)
	h.log.Info("random string", zap.String("str", str))
	e.Pong(c, e.Success, str)
}
