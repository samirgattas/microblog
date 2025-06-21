package user

import (
	"encoding/json"
	"io"
	"log/slog"
	"microblog/internal/core/domain"
	"microblog/internal/core/lib/customerror"
	userhandler "microblog/internal/core/port/handler/user"
	"microblog/internal/core/port/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Service user.UserService
}

func NewUserHandler(userService user.UserService) userhandler.UserHandler {
	return &userHandler{
		Service: userService,
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.ErrorContext(c, "read all error")
		c.Error(err)
		return
	}

	user := &domain.User{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		slog.ErrorContext(c, "user unmarshal error")
		c.Error(customerror.NewBadRequestError("invalid request body"))
		return
	}

	err = h.Service.Create(c, user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUser(c *gin.Context) {
user	IDStr := c.Param("user_id")
	if userIDStr == "" {
		slog.ErrorContext(c, "empty user_id")
		c.Error(customerror.NewBadRequestError("empty user_id"))
		return
	}

	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		slog.ErrorContext(c, "invalid user_id", slog.Any("user_id", userIDStr))
		c.Error(customerror.NewBadRequestError("invalid user_id"))
		return
	}

	user, err := h.Service.Get(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}
