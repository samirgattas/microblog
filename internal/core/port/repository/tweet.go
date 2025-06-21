package repository

import (
	"context"
	"microblog/internal/core/domain"
)

type TweetRepository interface {
	Save(context.Context, *domain.Tweet) error
	Get(context.Context, int64) (*domain.Tweet, error)
	Search(context.Context, domain.TweetSearchParams) (domain.TweetsSearchResult, error)
}
