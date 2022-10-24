package status

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/status"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Query(c *gin.Context)
	QueryByProblemVersion(c *gin.Context)
}

type handler struct {
	log           *zap.Logger
	statusService status.Service
	jwtService    global.JwtService
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:           g.Log,
		statusService: status.NewService(g),
		jwtService:    g.Jwt,
	}

	group.POST("/query", app.Query)
	group.POST("/query/problem_version", app.jwtService.Handler(true), app.QueryByProblemVersion)
}
