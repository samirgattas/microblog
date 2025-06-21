package user

import (
	"context"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/repository"
	"github.com/samirgattas/microblog/internal/core/port/service/user"
)

type userService struct {
	Repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) user.UserService {
	return &userService{
		Repository: userRepository,
	}
}

func (u *userService) Create(ctx context.Context, user *domain.User) error {
	err := u.Repository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Get(ctx context.Context, userID int64) (*domain.User, error) {
	user, err := u.Repository.Get(ctx, userID)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
