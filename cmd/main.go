package main

import (
	"microblog/internal/adapter/handler/healthcheck"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create HealthCheckHandler
	healthCheckHandler := healthcheck.NewHealthCheckHandler()

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	// Routes
	router.GET("/ping", healthCheckHandler.HealthCheck)

	// Run server
	router.Run("localhost:8080")
}
