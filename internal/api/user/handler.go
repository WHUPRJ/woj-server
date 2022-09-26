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
	userService user.Service
	jwtService  global.JwtService
}

func RouteRegister(g *global.Global, group *gin.RouterGroup) {
	app := &handler{
		log:         g.Log,
		userService: user.NewService(g),
		jwtService:  g.Jwt,
	}

	group.POST("/login", app.Login)
	group.POST("/create", app.Create)
	group.POST("/logout", app.jwtService.Handler(), app.Logout)
	group.POST("/profile", app.jwtService.Handler(), app.Profile)
}
