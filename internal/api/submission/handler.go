package submission

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/problem"
	"github.com/WHUPRJ/woj-server/internal/service/status"
	"github.com/WHUPRJ/woj-server/internal/service/submission"
	"github.com/WHUPRJ/woj-server/internal/service/task"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
	Rejudge(c *gin.Context)
}

type handler struct {
	log               *zap.Logger
	jwtService        global.JwtService
	problemService    problem.Service
	statusService     status.Service
	submissionService submission.Service
	taskService       task.Service
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:               g.Log,
		jwtService:        g.Jwt,
		problemService:    problem.NewService(g),
		statusService:     status.NewService(g),
		submissionService: submission.NewService(g),
		taskService:       task.NewService(g),
	}

	group.POST("/create", app.jwtService.Handler(true), app.Create)
	group.POST("/query", app.Query)
	group.POST("/rejudge", app.jwtService.Handler(true), app.Rejudge)
}
