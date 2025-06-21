package repository

import (
	"context"
	"microblog/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, userID int64) (*domain.User, error)
}
