package tweet

import (
	"context"
	"errors"
	"log/slog"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/lib/customerror"
	"github.com/samirgattas/microblog/internal/core/port/repository"
	"github.com/samirgattas/microblog/internal/core/port/service/tweet"
)

type tweetService struct {
	repository         repository.TweetRepository
	userRepository     repository.UserRepository
	followedRepository repository.FollowedRepository
}

func NewTweetService(tweetRepository repository.TweetRepository, userRepository repository.UserRepository, followedRepository repository.FollowedRepository) tweet.TweetService {
	return &tweetService{
		repository:         tweetRepository,
		userRepository:     userRepository,
		followedRepository: followedRepository,
	}
}

func (s *tweetService) Create(ctx context.Context, tweet *domain.Tweet) error {
	// Check if the post does not exceeds the maximum length
	if len(tweet.Post) > domain.MaxPostLength {
		slog.ErrorContext(ctx, "post too long", slog.Any("post_length", len(tweet.Post)))
		return customerror.NewBadRequestError("post too long")
	}

	// Check if the user exists
	_, err := s.userRepository.Get(ctx, tweet.UserID)
	if err != nil {
		var cErr customerror.NotFoundError
		// If the user does not exist, then return a bad request
		if errors.As(err, &cErr) {
			return customerror.NewBadRequestError("invalid user_id")
		}

		return err
	}

	// Save the tweet
	err = s.repository.Save(ctx, tweet)
	if err != nil {
		return err
	}

	return nil
}

func (s *tweetService) Get(ctx context.Context, tweetID int64) (*domain.Tweet, error) {
	// Get the tweet
	tweet, err := s.repository.Get(ctx, tweetID)
	if err != nil {
		return &domain.Tweet{}, err
	}

	return tweet, err
}

func (s *tweetService) Search(ctx context.Context, followerUserID int64, limit int64, offset int64) (domain.TweetsSearchResult, error) {
	followed, err := s.followedRepository.SearchByUserIDAndFollowedUserID(ctx, &followerUserID, nil)
	if err != nil {
		return domain.TweetsSearchResult{}, err
	}

	followedUserIDs := []int64{}
	for _, f := range followed {
		slog.InfoContext(ctx, "followed user_id", slog.Any("followed_user_id", f.FollowedUserID))
		followedUserIDs = append(followedUserIDs, f.FollowedUserID)
	}

	if limit == 0 || limit > domain.MaxLimit {
		limit = domain.MaxLimit
	}
	params := domain.TweetSearchParams{
		Limit:   limit,
		Offset:  offset,
		UserIDs: followedUserIDs,
	}

	result, err := s.repository.Search(ctx, params)
	if err != nil {
		return domain.TweetsSearchResult{}, err
	}

	return result, nil
}
