package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz handler godoc
// @Tags status
// @Summary Server status
// @Description This endpoint allow to see the serves status if 200 is received everything is working OK
// @Produce json
// @Success 200
// @Router /healthz [get]
func GETHealthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "online")
	}
}
