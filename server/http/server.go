package http

import (
	"fmt"

	"github.com/cjodra14/telemetry_backend/services"
	commonsRestConfig "github.com/cjodra14/telemetry_commons/configuration/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(config commonsRestConfig.HTTPServerConfiguration, telemetryService services.TelemetryService) error {
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

	telemetryGroup := router.Group("/")
	TelemetryRouter(telemetryGroup, telemetryService)

	infoGroup := router.Group("/")
	InfoRouter(infoGroup)

	serverAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)

	return router.Run(serverAddress)
}
