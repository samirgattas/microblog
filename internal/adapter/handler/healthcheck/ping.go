package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/internal/core/port/handler"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() handler.HealthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
