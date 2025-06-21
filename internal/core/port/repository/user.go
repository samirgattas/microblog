package repository

import (
	"context"
	"microblog/internal/core/domain"
)

type UserRepository interface {
	Save(context.Context, *domain.User) error
	Get(context.Context, int64) (*domain.User, error)
}
