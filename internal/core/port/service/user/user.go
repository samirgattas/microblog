package user

import (
	"context"
	"microblog/internal/core/domain"
)

type UserService interface {
	Create(context.Context, *domain.User) error
	Get(context.Context, int64) (*domain.User, error)
}
