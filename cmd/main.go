package main

import (
	"github.com/gin-gonic/gin"
	followedhandler "github.com/samirgattas/microblog/internal/adapter/handler/followed"
	"github.com/samirgattas/microblog/internal/adapter/handler/healthcheck"
	"github.com/samirgattas/microblog/internal/adapter/handler/middleware"
	tweethandler "github.com/samirgattas/microblog/internal/adapter/handler/tweet"
	userhandler "github.com/samirgattas/microblog/internal/adapter/handler/user"
	"github.com/samirgattas/microblog/internal/adapter/repository/followed"
	"github.com/samirgattas/microblog/internal/adapter/repository/tweet"
	"github.com/samirgattas/microblog/internal/adapter/repository/user"
	"github.com/samirgattas/microblog/internal/core/domain"
	followedservice "github.com/samirgattas/microblog/internal/core/service/followed"
	tweetservice "github.com/samirgattas/microblog/internal/core/service/tweet"
	userservice "github.com/samirgattas/microblog/internal/core/service/user"
)

func main() {
	userDB := inmemorystorage.NewStore()
	c := &config.Config{}
	c = c.NewConfig(userDB)
	handler := Container(c)

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

	// Tweet
	router.POST("/tweets", tweetHandler.CreateTweet)
	router.GET("/tweets/:tweet_id", tweetHandler.GetTweet)
	router.GET("/tweets", tweetHandler.SearchTweets)

	// Run server
	router.Run("localhost:8080")
}
