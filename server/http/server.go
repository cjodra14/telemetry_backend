package http

import (
	"fmt"
	"net/http"

	commonsRestConfig "github.com/cjodra14/telemetry_commons/configuration/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	TelemetryRouter *TelemetryRouter
}


func (server *Server) Init(config commonsRestConfig.HTTPServerConfiguration) error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/healthz"},
	}))

	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"*"}
	// corsConfig.AllowMethods = []string{"OPTIONS", "POST", "GET"}
	// corsConfig.AllowHeaders = []string{"Authentication", "Authorization", "content-type"}
	// corsConfig.ExposeHeaders = []string{"Authentication", "Authorization", "content-type"}
	router.Use(cors.New(corsConfig))
	router.GET("/healthz")

	router.GET("/doc/", func(c *gin.Context) { c.Redirect(http.StatusSeeOther, "/doc-api/index.html") })
	router.GET("/doc-api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "telemetry_backend"
	}))

	telmetryGroup := router.Group("/")
	server.TelemetryRouter.Router(telmetryGroup)

	serverAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)

	return router.Run(serverAddress)
}
