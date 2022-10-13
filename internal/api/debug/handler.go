package debug

import (
	"github.com/WHUPRJ/woj-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	randomString(c *gin.Context)
}

type handler struct {
	log *zap.Logger
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{g.Log}
	group.GET("/random", app.randomString)
}
