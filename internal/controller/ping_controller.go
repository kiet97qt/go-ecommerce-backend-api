package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-ecommerce-backend-api/internal/service"
)

// PingController coordinates ping related requests.
type PingController struct {
	pingService service.PingService
}

// NewPingController wires a ping controller with its dependencies.
func NewPingController(pingService service.PingService) *PingController {
	return &PingController{pingService: pingService}
}

// Ping responds with a pong payload to indicate service health.
func (pc *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": pc.pingService.Ping(ctx.Request.Context()),
	})
}
