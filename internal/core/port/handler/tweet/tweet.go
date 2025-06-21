package tweet

import "github.com/gin-gonic/gin"

type TweetHandler interface {
	CreateTweet(c *gin.Context)
	GetTweet(c *gin.Context)
	SearchTweets(c *gin.Context)
}
