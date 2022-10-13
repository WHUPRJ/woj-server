package problem

import (
	"github.com/WHUPRJ/woj-server/global"
	"github.com/WHUPRJ/woj-server/internal/service/problem"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Update(c *gin.Context)
	Search(c *gin.Context)
}

type handler struct {
	log            *zap.Logger
	problemService problem.Service
	jwtService     global.JwtService
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:            g.Log,
		problemService: problem.NewService(g),
		jwtService:     g.Jwt,
	}

	group.POST("/search", app.Search)
	group.POST("/update", app.jwtService.Handler(), app.Update)
}
