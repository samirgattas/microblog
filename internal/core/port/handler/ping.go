package handler

import "github.com/gin-gonic/gin"

type HealthCheckHandler interface {
	HealthCheck(*gin.Context)
}