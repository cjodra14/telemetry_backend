package http

import (
	"github.com/cjodra14/telemetry-iot/telemetry_backend/handlers"
	"github.com/cjodra14/telemetry-iot/telemetry_backend/services"
	"github.com/gin-gonic/gin"
)

type TelemetryRouter struct {
	telemetryService services.TelemetryService
}

func NewTelemetryRouter(telemetryService services.TelemetryService) *TelemetryRouter {
	return &TelemetryRouter{
		telemetryService: telemetryService,
	}
}

func (telemetryRouter TelemetryRouter) Router(router *gin.RouterGroup) {
	router.POST("/save", handlers.SaveTelemetry(telemetryRouter.telemetryService))
	router.GET("/user/:user-id", handlers.GetUserTelemetries(telemetryRouter.telemetryService))
}
