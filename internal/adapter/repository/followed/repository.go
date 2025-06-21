package followed

import (
	"context"
	"time"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/lib/customerror"
	"github.com/samirgattas/microblog/internal/core/port/repository"
)

var followedID = int64(0) // This is the entity id

type followedRepository struct {
	followedDB map[int64]domain.Followed
}

func NewFollowedRepository(followedDB map[int64]domain.Followed) repository.FollowedRepository {
	return &followedRepository{
		followedDB: followedDB,
	}
}

func (r *followedRepository) Save(ctx context.Context, followed *domain.Followed) error {
	now := time.Now()
	followed.CreatedAt = &now
	followed.UpdatedAt = &now
	followedID += 1
	followed.ID = followedID
	r.followedDB[followedID] = *followed
	return nil
}

func (r *followedRepository) Get(ctx context.Context, ID int64) (*domain.Followed, error) {
	followed, ok := r.followedDB[ID]
	// Check if followed exists
	if !ok {
		return &domain.Followed{}, customerror.NewNotFoundError("followed")
	}
	return &followed, nil
}

func (r *followedRepository) Update(ctx context.Context, followed *domain.Followed) error {
	now := time.Now()
	followed.UpdatedAt = &now
	if _, ok := r.followedDB[followed.UserID]; !ok {
		return customerror.NewNotFoundError("followed")
	}
	r.followedDB[followed.UserID] = *followed
	return nil
}

func (r *followedRepository) SearchByUserIDAndFollowedUserID(ctx context.Context, followerUserID *int64, followedUserID *int64) ([]domain.Followed, error) {
	followedArray := []domain.Followed{}
	for _, followed := range r.followedDB {
		// Discad followed because is not the right follower_user_id
		if followerUserID != nil && followed.UserID != *followerUserID {
			continue
		}
		// Discad followed because is not the right followed_user_id
		if followedUserID != nil && followed.FollowedUserID != *followedUserID {
			continue
		}
		followedArray = append(followedArray, followed)
	}
	return followedArray, nil
}
