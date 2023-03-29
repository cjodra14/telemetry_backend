package handlers

import (
	"net/http"

	"github.com/cjodra14/telemetry_backend/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GETTelemetry handler godoc
//
//	@Tags			telemetry
//	@Summary		This endpoint allows the client to get user telemetries
//	@Description	Client get the telemetry and it is saved on it's user on the database
//	@Param			user-id	path	string	true	"User ID"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/user/{user-id} [get]
func GetUserTelemetries(telemetryService services.TelemetryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user-id")

		telemetries, err := telemetryService.GetuserTelemetries(c, userID)
		if err != nil {
			if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
				logrus.Debug(err)
			}

			return
		}

		c.JSON(http.StatusOK, telemetries)
	}
}
