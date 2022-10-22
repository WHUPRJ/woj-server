package problem

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/problem"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Details(c *gin.Context)
	Search(c *gin.Context)
	Update(c *gin.Context)
}

type handler struct {
	log            *zap.Logger
	jwtService     global.JwtService
	problemService problem.Service
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:            g.Log,
		jwtService:     g.Jwt,
		problemService: problem.NewService(g),
	}

	group.POST("/details", app.jwtService.Handler(false), app.Details)
	group.POST("/search", app.Search)
	group.POST("/update", app.jwtService.Handler(true), app.Update)
}
