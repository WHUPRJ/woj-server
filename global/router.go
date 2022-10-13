package global

import (
	"github.com/gin-gonic/gin"
)

type EndpointInfo struct {
	Version  string
	Path     string
	Register func(g *Global, group *gin.RouterGroup)
}
