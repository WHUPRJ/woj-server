package user

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/repo/model"
	"github.com/WHUPRJ/woj-server/internal/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	// List(c *gin.Context)

	tokenNext(c *gin.Context, user *model.User)
}

type handler struct {
	log         *zap.Logger
	userService user.Service
	jwtService  global.JwtService
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:         g.Log,
		userService: user.NewUserService(g),
		jwtService:  g.Jwt,
	}

	group.POST("/login", app.Login)
	group.POST("/create", app.jwtService.Handler(), app.Create)
	// group.GET("/", app.List)
}
