package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/config"
	"github.com/samirgattas/microblog/internal/adapter/handler/middleware"
	inmemorystorage "github.com/samirgattas/microblog/internal/core/lib/customerror/in_memory_storage"
	followedhandlerr "github.com/samirgattas/microblog/internal/core/port/handler/followed"
	healthcheckhandlerr "github.com/samirgattas/microblog/internal/core/port/handler/healthcheck"
	tweethandlerr "github.com/samirgattas/microblog/internal/core/port/handler/tweet"
	userhandlerr "github.com/samirgattas/microblog/internal/core/port/handler/user"
)

type Handler struct {
	healthCheckHandler healthcheckhandlerr.HealthCheckHandler
	userHandler        userhandlerr.UserHandler
	followedHandler    followedhandlerr.FollowedHandler
	tweetHandler       tweethandlerr.TweetHandler
}

func main() {
	userDB := inmemorystorage.NewStore()
	c := &config.Config{}
	c = c.NewConfig(userDB)
	handler := Container(c)

	router := gin.Default()

	Routes(router, handler)
	
	// Run server
	router.Run("localhost:8080")
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
