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
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Profile(c *gin.Context)
}

type handler struct {
	log         *zap.Logger
	jwtService  global.JwtService
	userService user.Service
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:         g.Log,
		jwtService:  g.Jwt,
		userService: user.NewService(g),
	}

	group.POST("/create", app.Create)
	group.POST("/login", app.Login)
	group.POST("/logout", app.jwtService.Handler(true), app.Logout)
	group.POST("/profile", app.jwtService.Handler(true), app.Profile)
}
