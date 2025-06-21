package followed

import (
	"context"
	"microblog/internal/core/domain"
)

type FollowedService interface {
	Create(ctx context.Context, followed *domain.Followed) error
	Get(ctx context.Context, followedID int64) (*domain.Followed, error)
	Update(ctx context.Context, followedID int64, followedPatchCmd *domain.FollowedPatchCommand) (*domain.Followed, error)
	Search(ctx context.Context, followerUserID *int64, followedUserID *int64) ([]domain.Followed, error)
}
