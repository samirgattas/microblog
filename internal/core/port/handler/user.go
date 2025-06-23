package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(*gin.Context)
	GetUser(*gin.Context)
}
