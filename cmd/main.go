package main

import (
	"microblog/internal/adapter/handler/healthcheck"
	"microblog/internal/adapter/handler/middleware"
	userhandler "microblog/internal/adapter/handler/user"
	"microblog/internal/adapter/repository/user"
	"microblog/internal/core/domain"
	userservice "microblog/internal/core/service/user"

	"github.com/gin-gonic/gin"
)

func main() {
	usersDB := make(map[int64]domain.User)
	// Create User repository
	userRepository := user.NewUserRepository(usersDB)
	// Create User service
	userService := userservice.NewUserService(userRepository)
	// Create User handler
	userHandler := userhandler.NewUserHandler(userService)

	// Create HealthCheckHandler
	healthCheckHandler := healthcheck.NewHealthCheckHandler()

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	// Routes
	router.GET("/ping", healthCheckHandler.HealthCheck)

	// Run server
	router.Run("localhost:8080")
}
