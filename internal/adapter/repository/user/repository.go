package user

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/lib/customerror"
	inmemorystore "github.com/samirgattas/microblog/internal/core/lib/customerror/in_memory_store"
	"github.com/samirgattas/microblog/internal/core/port/repository"
)

type userRepository struct {
	usersDB inmemorystore.Store
}

func NewUserRepository(usersDB inmemorystore.Store) repository.UserRepository {
	return &userRepository{usersDB: usersDB}
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) error {
	now := time.Now()
	user.CreatedAt = &now
	u.usersDB.SaveWithID(user.ID, *user)
	return nil
}

func (u *userRepository) Get(ctx context.Context, userID int64) (*domain.User, error) {
	item, err := u.usersDB.Get(userID)
	if err != nil {
		if errors.Is(err, inmemorystore.ErrNotFound) {
			slog.ErrorContext(ctx, "user not found", slog.Any("user_id", userID))
			return &domain.User{}, customerror.NewNotFoundError("user")
		}
		return &domain.User{}, err
	}
	user, ok := item.(domain.User)
	if !ok {
		slog.WarnContext(ctx, "invalid item", slog.Any("entity", "userDB"), slog.Any("id", userID))
	}
	return &user, nil
}
