package tweet

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/handler"
	"github.com/samirgattas/microblog/internal/core/port/service"
	"github.com/samirgattas/microblog/lib/customerror"
)

type tweetHandler struct {
	service service.TweetService
}

func NewTweetHandler(tweetService service.TweetService) handler.TweetHandler {
	return &tweetHandler{service: tweetService}
}

func (h *tweetHandler) CreateTweet(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.Error(customerror.NewBadRequestError("invalid format"))
		return
	}

	tweet := &domain.Tweet{}
	err = json.Unmarshal(jsonData, &tweet)
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.Error(customerror.NewBadRequestError("invalid request body"))
		return
	}

	err = h.service.Create(c, tweet)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, tweet)
}

func (h *tweetHandler) GetTweet(c *gin.Context) {
	tweetIDStr := c.Param("tweet_id")
	if tweetIDStr == "" {
		slog.ErrorContext(c, "empty tweet_id")
		c.Error(customerror.NewBadRequestError("empty tweet_id"))
		return
	}

	tweetID, err := strconv.ParseInt(tweetIDStr, 10, 64)
	if err != nil {
		slog.ErrorContext(c, err.Error(), slog.Any("tweet_id", tweetIDStr))
		c.Error(customerror.NewBadRequestError("invalid tweet_id"))
		return
	}

	tweet, err := h.service.Get(c, tweetID)
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tweet)
}

func (h *tweetHandler) SearchTweets(c *gin.Context) {
	var err error

	userID := int64(0)
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		slog.ErrorContext(c, "user_id is empty")
		c.Error(customerror.NewBadRequestError("empty user_id"))
		return
	}
	userID, err = strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		slog.ErrorContext(c, err.Error(), slog.Any("user_id", userIDStr))
		c.Error(customerror.NewBadRequestError("invalid user_id"))
		return
	}

	limit := int64(0)
	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			slog.ErrorContext(c, err.Error(), slog.Any("limit", limitStr))
			c.Error(customerror.NewBadRequestError("invalid limit"))
			return
		}
	}

	offset := int64(0)
	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			slog.ErrorContext(c, err.Error(), slog.Any("offset", offsetStr))
			c.Error(customerror.NewBadRequestError("invalid offset"))
			return
		}
	}

	tweet, err := h.service.Search(c, userID, limit, offset)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tweet)
}
