package router

import (
	"github.com/WHUPRJ/woj-server/internal/controller/debug"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

// @title OJ Server API Documentation
// @version 1.0
// @BasePath /api
func setupApi(g *global.Global, root *gin.RouterGroup) {
	for _, v := range endpoints {
		group := root.Group(v.Version).Group(v.Path)
		v.Register(g, group)
	}
}

var endpoints = []model.EndpointInfo{
	{Version: "", Path: "/debug", Register: debug.RouteRegister},
}
