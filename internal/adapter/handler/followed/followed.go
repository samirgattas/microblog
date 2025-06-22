package followed

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/lib/customerror"
	"github.com/samirgattas/microblog/internal/core/port/handler/followed"
	followedsrv "github.com/samirgattas/microblog/internal/core/port/service/followed"
)

type followedHandler struct {
	service followedsrv.FollowedService
}

func NewFollowedHandler(followedService followedsrv.FollowedService) followed.FollowedHandler {
	return &followedHandler{
		service: followedService,
	}
}

func (h *followedHandler) CreateFollowed(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.ErrorContext(c, "read all error")
		c.Error(err)
		return
	}

	followed := &domain.Followed{}
	err = json.Unmarshal(jsonData, &followed)
	if err != nil {
		slog.ErrorContext(c, "followed unmarshal error")
		c.Error(customerror.NewBadRequestError("invalid request body"))
		return
	}

	err = h.service.Create(c, followed)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, followed)
}

func (h *followedHandler) GetFollowed(c *gin.Context) {
	IDStr := c.Param("followed_id")
	if IDStr == "" {
		slog.ErrorContext(c, "empty followed_id")
		c.Error(customerror.NewBadRequestError("empty followed_id"))
		return
	}

	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		slog.ErrorContext(c, fmt.Sprintf("invalid followed_id: %s", IDStr))
		c.Error(customerror.NewBadRequestError("invalid followed_id"))
		return
	}
	user, err := h.service.Get(c, ID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *followedHandler) UpdateFollowed(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.ErrorContext(c, "read all error")
		c.Error(err)
		return
	}

	followedIDStr := c.Param("followed_id")
	if followedIDStr == "" {
		slog.ErrorContext(c, "empty followed_id")
		c.Error(customerror.NewBadRequestError("empty followed_id"))
		return
	}
	followedID, err := strconv.ParseInt(followedIDStr, 10, 64)
	if err != nil {
		slog.ErrorContext(c, fmt.Sprintf("invalid followed_id: %s", followedIDStr))
		c.Error(customerror.NewBadRequestError("invalid followed_id"))
		return
	}

	followedPatchCmd := &domain.FollowedPatchCommand{}
	err = json.Unmarshal(jsonData, &followedPatchCmd)
	if err != nil {
		slog.ErrorContext(c, "followed unmarshal error")
		c.Error(customerror.NewBadRequestError("invalid request body"))
		return
	}

	followed, err := h.service.Update(c, followedID, followedPatchCmd)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, followed)
}

func (h *followedHandler) SearchFollowed(c *gin.Context) {
	var userID *int64
	userIDStr := c.Query("user_id")
	if userIDStr != "" {
		userIDTmp, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			slog.ErrorContext(c, fmt.Sprintf("invalid user_id: %s", userIDStr))
			c.Error(customerror.NewBadRequestError("invalid user_id"))
			return
		}
		userID = &userIDTmp
	}

	var followedUserID *int64
	followedUserIDStr := c.Query("followed_user_id")
	if followedUserIDStr != "" {
		followedUserIDTmp, err := strconv.ParseInt(followedUserIDStr, 10, 64)
		if err != nil {
			slog.ErrorContext(c, fmt.Sprintf("invalid user_id: %s", followedUserIDStr))
			c.Error(customerror.NewBadRequestError("invalid user_id"))
			return
		}
		followedUserID = &followedUserIDTmp
	}

	if userID == nil && followedUserID == nil {
		c.Error(customerror.NewBadRequestError("invalid user_id"))
		return
	}

	followed, err := h.service.Search(c, userID, followedUserID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": len(followed), "followed": followed})
}
