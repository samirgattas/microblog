package healthcheck

import (
	"microblog/internal/core/port/handler/healthcheck"
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() healthcheck.HealthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
