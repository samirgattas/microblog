package main

import (
	"github.com/samirgattas/microblog/config"
	followedhandler "github.com/samirgattas/microblog/internal/adapter/handler/followed"
	"github.com/samirgattas/microblog/internal/adapter/handler/healthcheck"
	tweethandler "github.com/samirgattas/microblog/internal/adapter/handler/tweet"
	userhandler "github.com/samirgattas/microblog/internal/adapter/handler/user"
	"github.com/samirgattas/microblog/internal/adapter/repository/followed"
	"github.com/samirgattas/microblog/internal/adapter/repository/tweet"
	"github.com/samirgattas/microblog/internal/adapter/repository/user"
	followedservice "github.com/samirgattas/microblog/internal/core/service/followed"
	tweetservice "github.com/samirgattas/microblog/internal/core/service/tweet"
	userservice "github.com/samirgattas/microblog/internal/core/service/user"
)

func Container(c *config.Config) Handler {
	// Create User repository
	userRepository := user.NewUserRepository(c.UserDB)
	// Create User service
	userService := userservice.NewUserService(userRepository)
	// Create User handler
	userHandler := userhandler.NewUserHandler(userService)

	// Create Followed repository
	followedRepository := followed.NewFollowedRepository(c.FollowedDB)
	// Create Followed service
	followedService := followedservice.NewFollowedService(followedRepository, userRepository)
	// Create Followed handler
	followedHandler := followedhandler.NewFollowedHandler(followedService)

	// Create Tweet repository
	tweetRepository := tweet.NewTweetRepository(c.TweetDB)
	// Create Tweet service
	tweetService := tweetservice.NewTweetService(tweetRepository, userRepository, followedRepository)
	// Create Tweet handler
	tweetHandler := tweethandler.NewTweetHandler(tweetService)

	// Create HealthCheckHandler
	healthCheckHandler := healthcheck.NewHealthCheckHandler()

	handler := Handler{
		healthCheckHandler: healthCheckHandler,
		userHandler:        userHandler,
		followedHandler:    followedHandler,
		tweetHandler:       tweetHandler,
	}

	return handler
}
