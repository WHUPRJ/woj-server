package router

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	_ "github.com/WHUPRJ/woj-server/internal/router/docs"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

func InitRouters(g *global.Global) *gin.Engine {
	gin.SetMode(utils.If(g.Conf.Development, gin.DebugMode, gin.ReleaseMode).(string))

	r := gin.New()
	r.MaxMultipartMemory = 8 << 20

	// Logger middleware and debug
	if g.Conf.Development {
		// Gin's default logger is pretty enough
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
		// add prof
		pprof.Register(r)
	} else {
		r.Use(ginZap.Ginzap(g.Log, time.RFC3339, false))
		r.Use(ginZap.RecoveryWithZap(g.Log, true))
	}

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	// Prometheus middleware
	r.Use(g.Stat.Handler())

	// metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health
	r.GET("/health", func(c *gin.Context) {
		resp := &struct {
			Timestamp time.Time `json:"timestamp"`
			Status    string    `json:"status"`
		}{
			Timestamp: time.Now(),
			Status:    "ok",
		}
		c.JSON(http.StatusOK, resp)
	})

	// api
	api := r.Group("/api/")
	setupApi(g, api)

	// static files
	r.Static("/static", "./resource/frontend/static")
	r.StaticFile("/", "./resource/frontend/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("./resource/frontend/index.html")
	})

	return r
}
