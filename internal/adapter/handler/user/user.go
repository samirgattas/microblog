package user

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

type userHandler struct {
	Service service.UserService
}

func NewUserHandler(userService service.UserService) handler.UserHandler {
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

	c.JSON(http.StatusCreated, user)
}

func (h *userHandler) GetUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		slog.ErrorContext(c, "empty user_id")
		c.Error(customerror.NewBadRequestError("empty user_id"))
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
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
