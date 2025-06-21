package followed

import (
	"context"
	"log/slog"
	"microblog/internal/core/domain"
	"microblog/internal/core/lib/customerror"
	"microblog/internal/core/port/repository"
	"microblog/internal/core/port/service/followed"
)

type followedService struct {
	repository     repository.FollowedRepository
	userRepository repository.UserRepository
}

func NewFollowedService(followedRepository repository.FollowedRepository, userRepository repository.UserRepository) followed.FollowedService {
	return &followedService{
		repository:     followedRepository,
		userRepository: userRepository,
	}
}

func (s *followedService) Create(ctx context.Context, followed *domain.Followed) error {
	// Check if the user id exists
	_, err := s.userRepository.Get(ctx, followed.UserID)
	if err != nil {
		return err
	}

	// Check if the followed user id exists
	_, err = s.userRepository.Get(ctx, followed.FollowedUserID)
	if err != nil {
		return err
	}

	// Set enabled to true because the creation occurs when the user_id starts following the followed_user_id
	followed.Enabled = true

	// Save the followed entity
	err = s.repository.Save(ctx, followed)
	if err != nil {
		return err
	}
	return nil
}

func (s *followedService) Get(ctx context.Context, followedID int64) (*domain.Followed, error) {
	// Get the followed entity
	followed, err := s.repository.Get(ctx, followedID)
	if err != nil {
		return &domain.Followed{}, err
	}

	return followed, nil
}

func (s *followedService) Update(ctx context.Context, followedID int64, followedPatchCmd *domain.FollowedPatchCommand) (*domain.Followed, error) {
	// Check if enabled is set
	if followedPatchCmd.Enabled == nil {
		return &domain.Followed{}, customerror.NewBadRequestError("enabled should not be null")
	}

	// Get the followed entity
	followed, err := s.repository.Get(ctx, followedID)
	if err != nil {
		return &domain.Followed{}, err
	}

	// If the enabled field is not modified, then return
	if followed.Enabled == *followedPatchCmd.Enabled {
		return followed, nil
	}

	// Set the enabled field
	followed.Enabled = *followedPatchCmd.Enabled

	// Update followed
	err = s.repository.Update(ctx, followed)
	if err != nil {
		return &domain.Followed{}, err
	}

	return followed, nil
}

func (s *followedService) Search(ctx context.Context, userID *int64, followedUserID *int64) ([]domain.Followed, error) {
	// Get the followed entity
	followed, err := s.repository.SearchByUserIDAndFollowedUserID(ctx, userID, followedUserID)
	if err != nil {
		return []domain.Followed{}, err
	}

	// Followed does not exist
	if len(followed) == 0 {
		slog.ErrorContext(ctx, "empty followed result")
		return []domain.Followed{}, customerror.NewNotFoundError("followed")
	}

	return followed, nil
}
