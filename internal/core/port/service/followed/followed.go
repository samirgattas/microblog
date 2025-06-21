package followed

import (
	"context"
	"microblog/internal/core/domain"
)

type FollowedService interface {
	CreateFollowed(context.Context, *domain.Followed) error
	GetFollowed(context.Context, int64) (*domain.Followed, error)
	UpdateFollowed(context.Context, int64, *domain.FollowedPatchCommand) (*domain.Followed, error)
	SearchFollowed(context.Context, *int64, *int64) ([]domain.Followed, error)
}
