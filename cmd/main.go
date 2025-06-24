package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/config"
	"github.com/samirgattas/microblog/internal/adapter/handler/middleware"
	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/handler"
	inmemorystore "github.com/samirgattas/microblog/lib/in_memory_store"
)

type Handler struct {
	healthCheckHandler handler.HealthCheckHandler
	userHandler        handler.UserHandler
	followedHandler    handler.FollowedHandler
	tweetHandler       handler.TweetHandler
}

func main() {
	userDB := inmemorystore.NewStore()
	followedDB := make(map[int64]domain.Followed)
	tweetDB := make(map[int64]domain.Tweet)
	c := &config.Config{}
	c = c.NewConfig(userDB, followedDB, tweetDB)
	h := Container(c)

	router := gin.Default()

	Routes(router, h)

	// Run server
	if err := router.Run(c.Domain + ":8080"); err != nil {
		os.Exit(1)
	}
}

func Routes(router *gin.Engine, h Handler) {
	router.Use(middleware.ErrorHandler())

	// Routes
	// Health check
	router.GET("/ping", h.healthCheckHandler.HealthCheck)

	// User
	router.POST("/users", h.userHandler.CreateUser)
	router.GET("/users/:user_id", h.userHandler.GetUser)

	// Followed
	router.POST("/followed", h.followedHandler.CreateFollowed)
	router.GET("/followed/:followed_id", h.followedHandler.GetFollowed)
	router.PATCH("/followed/:followed_id", h.followedHandler.UpdateFollowed)
	router.GET("/followed", h.followedHandler.SearchFollowed)

	// Tweet
	router.POST("/tweets", h.tweetHandler.CreateTweet)
	router.GET("/tweets/:tweet_id", h.tweetHandler.GetTweet)
	router.GET("/tweets", h.tweetHandler.SearchTweets)
}
