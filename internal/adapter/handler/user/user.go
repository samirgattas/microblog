package user

import (
	"encoding/json"
	"fmt"
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

func (u *userHandler) CreateUser(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error("read all error")
		c.Error(err)
		return
	}
	user := &domain.User{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		slog.Error("user unmarshal error")
		c.Error(customerror.NewBadRequestError("invalid request body"))
		return
	}
	err = u.Service.Create(user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *userHandler) GetUser(c *gin.Context) {
	IDStr := c.Param("user_id")
	if IDStr == "" {
		slog.Error("empty user_id")
		c.Error(customerror.NewBadRequestError("empty user_id"))
		return
	}

	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		slog.Error(fmt.Sprintf("invalid user_id: %s", IDStr))
		c.Error(customerror.NewBadRequestError("invalid user_id"))
		return
	}
	user, err := u.Service.Get(int64(ID))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}
