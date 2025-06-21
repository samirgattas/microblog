package followed

import "github.com/gin-gonic/gin"

type FollowedHandler interface {
	CreateFollowed(*gin.Context)
	GetFollowed(*gin.Context)
	UpdateFollowed(*gin.Context)
	SearchFollowed(*gin.Context)
}
