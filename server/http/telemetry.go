package http

import (
	"net/http"

	"github.com/cjodra14/telemetry_backend/handlers"
	"github.com/cjodra14/telemetry_backend/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InfoRouter(router *gin.RouterGroup) {
	router.GET("/healthz")

	router.GET("/doc/", func(c *gin.Context) { c.Redirect(http.StatusSeeOther, "/doc-api/index.html") })
	router.GET("/doc-api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "telemetry_backend"
	}))
}

func TelemetryRouter(router *gin.RouterGroup, telemetryService services.TelemetryService) {
	router.POST("/save", handlers.SaveTelemetry(telemetryService))
	router.GET("/user/:user-id", handlers.GetUserTelemetries(telemetryService))
}
