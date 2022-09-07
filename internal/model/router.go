package model

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/gin-gonic/gin"
)

type EndpointInfo struct {
	Version  string
	Path     string
	Register func(g *global.Global, group *gin.RouterGroup)
}
