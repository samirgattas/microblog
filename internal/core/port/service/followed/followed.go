package followed

import (
	"context"
	"microblog/internal/core/domain"
)

type FollowedService interface {
	Create(context.Context, *domain.Followed) error
	Get(context.Context, int64) (*domain.Followed, error)
	Update(context.Context, int64, *domain.FollowedPatchCommand) (*domain.Followed, error)
	Search(context.Context, *int64, *int64) ([]domain.Followed, error)
}
