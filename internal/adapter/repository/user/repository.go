package user

import (
	"context"
	"log/slog"
	"microblog/internal/core/domain"
	"microblog/internal/core/lib/customerror"
	"microblog/internal/core/port/repository"
	"time"
)

type userRepository struct {
	usersDB map[int64]domain.User
}

func NewUserRepository(usersDB map[int64]domain.User) repository.UserRepository {
	return &userRepository{usersDB: usersDB}
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) error {
	now := time.Now()
	user.CreatedAt = &now
	u.usersDB[user.ID] = *user
	return nil
}

func (u *userRepository) Get(ctx context.Context, userID int64) (*domain.User, error) {
	user, ok := u.usersDB[userID]
	if !ok {
		slog.ErrorContext(ctx, "user not found", slog.Any("user_id", userID))
		return &domain.User{}, customerror.NewNotFoundError("user")
	}
	return &user, nil
}
