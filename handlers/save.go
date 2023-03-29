package handlers

import (
	"net/http"

	"github.com/cjodra14/telemetry_backend/api/models"
	"github.com/cjodra14/telemetry_backend/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// POSTTelemetry handler godoc
//
//	@Tags			telemetry
//	@Summary		This endpoint allows the client to post the telemetry
//	@Description	Client Post the telmetry and it is saved on it's user on the database
//	@Accept			json
//	@Param			telemetry	body	models.Telemetry	true	"Telemetry"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/save [post]
func SaveTelemetry(telemetryService services.TelemetryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		telemetry := models.Telemetry{}
		if err := c.BindJSON(&telemetry); err != nil {
			if err := c.AbortWithError(http.StatusBadRequest, err); err != nil {
				logrus.Debug(err)
			}

			return
		}

		if err := telemetryService.SaveTelemetry(c, telemetry); err != nil {
			if err := c.AbortWithError(http.StatusBadRequest, err); err != nil {
				logrus.Debug(err)
			}

			return
		}

		c.Status(http.StatusOK)
	}
}
