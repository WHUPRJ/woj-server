package user

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Create(c *gin.Context)
	// List(c *gin.Context)
}

type handler struct {
	log     *zap.Logger
	service user.Service
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:     g.Log,
		service: user.NewUserService(g),
	}

	group.POST("/", app.Create)
	// group.GET("/", app.List)
}
