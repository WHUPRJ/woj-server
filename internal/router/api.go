package router

import (
	"github.com/WHUPRJ/woj-server/internal/api/debug"
	"github.com/WHUPRJ/woj-server/internal/api/problem"
	"github.com/WHUPRJ/woj-server/internal/api/user"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/gin-gonic/gin"
)

// @title OJ Server API Documentation
// @version                     1.0
// @BasePath                    /api
// @securityDefinitions.apikey  Authentication
// @in                          header
// @name                        Authorization
func setupApi(g *global.Global, root *gin.RouterGroup) {
	for _, v := range endpoints {
		group := root.Group(v.Version).Group(v.Path)
		v.Register(g, group)
	}
}

var endpoints = []global.EndpointInfo{
	{Version: "", Path: "/debug", Register: debug.RouteRegister},
	{Version: "/v1", Path: "/user", Register: user.RouteRegister},
	{Version: "/v1", Path: "/problem", Register: problem.RouteRegister},
}
