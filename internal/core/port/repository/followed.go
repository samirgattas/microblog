package repository

import (
	"context"
	"microblog/internal/core/domain"
)

type FollowedRepository interface {
	Save(context.Context, *domain.Followed) error
	Get(context.Context, int64) (*domain.Followed, error)
	Update(context.Context, *domain.Followed) error
	SearchByUserIDAndFollowedUserID(context.Context, *int64, *int64) ([]domain.Followed, error)
}
