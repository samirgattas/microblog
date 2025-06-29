package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/config"
	followedhandler "github.com/samirgattas/microblog/internal/adapter/handler/followed"
	"github.com/samirgattas/microblog/internal/adapter/handler/healthcheck"
	"github.com/samirgattas/microblog/internal/adapter/handler/middleware"
	tweethandler "github.com/samirgattas/microblog/internal/adapter/handler/tweet"
	userhandler "github.com/samirgattas/microblog/internal/adapter/handler/user"
	"github.com/samirgattas/microblog/internal/adapter/repository/followed"
	"github.com/samirgattas/microblog/internal/adapter/repository/tweet"
	"github.com/samirgattas/microblog/internal/adapter/repository/user"
	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/handler"
	followedservice "github.com/samirgattas/microblog/internal/core/service/followed"
	tweetservice "github.com/samirgattas/microblog/internal/core/service/tweet"
	userservice "github.com/samirgattas/microblog/internal/core/service/user"
	"github.com/samirgattas/microblog/lib/database"
)

type Handler struct {
	healthCheckHandler handler.HealthCheckHandler
	userHandler        handler.UserHandler
	followedHandler    handler.FollowedHandler
	tweetHandler       handler.TweetHandler
}

func main() {
	c := &config.Config{}
	c = c.NewConfig()
	app := NewApp(*c)
	h := Container(app)

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

func Container(app App) Handler {
	// Create User repository
	userRepository := user.NewUserRepository(app.Database)
	// Create User service
	userService := userservice.NewUserService(userRepository)
	// Create User handler
	userHandler := userhandler.NewUserHandler(userService)

	// Create Followed repository
	followedRepository := followed.NewFollowedRepository(app.Database)
	// Create Followed service
	followedService := followedservice.NewFollowedService(followedRepository, userRepository)
	// Create Followed handler
	followedHandler := followedhandler.NewFollowedHandler(followedService)

	// Create Tweet repository
	tweetRepository := tweet.NewTweetRepository(app.TweetDB)
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

type App struct {
	Database *sql.DB
	TweetDB  map[int64]domain.Tweet
}

func NewApp(c config.Config) App {
	// Init MySQL
	db, err := database.NewMySQL(c.MySQL.User, c.MySQL.Password, c.MySQL.Domain, c.MySQL.Port)
	if err != nil {
		log.Fatalf("error init mysql")
	}

	tweetDB := make(map[int64]domain.Tweet)

	return App{
		Database: db,
		TweetDB:  tweetDB,
	}
}
