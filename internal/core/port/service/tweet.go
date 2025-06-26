package service

import (
	"context"

	"github.com/samirgattas/microblog/internal/core/domain"
)

type TweetService interface {
	Create(ctx context.Context, tweet *domain.Tweet) error
	Get(ctx context.Context, tweetID int64) (*domain.Tweet, error)
	Search(ctx context.Context, followerUserID int64, limit int64, offset int64) (domain.TweetsSearchResult, error)
}
