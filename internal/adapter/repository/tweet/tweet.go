package tweet

import (
	"cmp"
	"context"
	"log/slog"
	"slices"
	"time"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/repository"
	"github.com/samirgattas/microblog/lib/customerror"
)

var tweetID = int64(0) // This is the entity id

type tweetRepository struct {
	tweetDB map[int64]domain.Tweet
}

func NewTweetRepository(tweetDB map[int64]domain.Tweet) repository.TweetRepository {
	return &tweetRepository{
		tweetDB: tweetDB,
	}
}

func (r *tweetRepository) Save(ctx context.Context, tweet *domain.Tweet) error {
	now := time.Now()
	tweetID++
	tweet.ID = tweetID
	tweet.CreatedAt = &now
	r.tweetDB[tweetID] = *tweet
	return nil
}

func (r *tweetRepository) Get(ctx context.Context, tweetID int64) (*domain.Tweet, error) {
	tweet, ok := r.tweetDB[tweetID]
	if !ok {
		slog.ErrorContext(ctx, "tweet not found", slog.Any("tweet_id", tweetID))
		return &domain.Tweet{}, customerror.NewNotFoundError("tweet")
	}
	return &tweet, nil
}

func (r *tweetRepository) Search(ctx context.Context, params domain.TweetSearchParams) (domain.TweetsSearchResult, error) {
	tweets := []domain.Tweet{}
	for _, tweet := range r.tweetDB {
		if !slices.Contains(params.UserIDs, tweet.UserID) {
			continue
		}
		tweets = append(tweets, tweet)
	}

	paging := domain.Paging{
		Total:  int64(len(tweets)),
		Limit:  params.Limit,
		Offset: params.Offset,
	}
	slices.SortFunc(tweets, sortByDescID)

	bottom := paging.Offset
	top := paging.Offset + paging.Limit
	if top > int64(len(tweets)) {
		top = int64(len(tweets))
	}
	result := domain.TweetsSearchResult{
		Results: tweets[bottom:top],
		Paging:  paging,
	}
	return result, nil
}

func sortByDescID(t1, t2 domain.Tweet) int {
	return cmp.Compare(t2.ID, t1.ID)
}
