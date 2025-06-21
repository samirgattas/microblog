package main

import (
	followedhandler "microblog/internal/adapter/handler/followed"
	"microblog/internal/adapter/handler/healthcheck"
	"microblog/internal/adapter/handler/middleware"
	tweethandler "microblog/internal/adapter/handler/tweet"
	userhandler "microblog/internal/adapter/handler/user"
	"microblog/internal/adapter/repository/followed"
	"microblog/internal/adapter/repository/tweet"
	"microblog/internal/adapter/repository/user"
	"microblog/internal/core/domain"
	followedservice "microblog/internal/core/service/followed"
	tweetservice "microblog/internal/core/service/tweet"
	userservice "microblog/internal/core/service/user"

	"github.com/gin-gonic/gin"
)

func main() {
	usersDB := make(map[int64]domain.User)
	followedDB := make(map[int64]domain.Followed)
	tweetDB := make(map[int64]domain.Tweet)

	// Create User repository
	userRepository := user.NewUserRepository(usersDB)
	// Create User service
	userService := userservice.NewUserService(userRepository)
	// Create User handler
	userHandler := userhandler.NewUserHandler(userService)

	// Create Followed repository
	followedRepository := followed.NewFollowedRepository(followedDB)
	// Create Followed service
	followedService := followedservice.NewFollowedService(followedRepository, userRepository)
	// Create Followed handler
	followedHandler := followedhandler.NewFollowedHandler(followedService)

	// Create Tweet repository
	tweetRepository := tweet.NewTweetRepository(tweetDB)
	// Create Tweet service
	tweetService := tweetservice.NewTweetService(tweetRepository, userRepository, followedRepository)
	// Create Tweet handler
	tweetHandler := tweethandler.NewTweetHandler(tweetService)

	// Create HealthCheckHandler
	healthCheckHandler := healthcheck.NewHealthCheckHandler()

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	// Routes
	// Health check
	router.GET("/ping", healthCheckHandler.HealthCheck)

	// User
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:user_id", userHandler.GetUser)

	// Followed
	router.POST("/followed", followedHandler.CreateFollowed)
	router.GET("/followed/:followed_id", followedHandler.GetFollowed)
	router.PATCH("/followed/:followed_id", followedHandler.UpdateFollowed)
	router.GET("/followed", followedHandler.SearchFollowed)

	// Run server
	router.Run("localhost:8080")
}
