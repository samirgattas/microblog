package user

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/repository"
	"github.com/samirgattas/microblog/lib/customerror"
	inmemorystore "github.com/samirgattas/microblog/lib/in_memory_store"
)

type userRepository struct {
	usersDB inmemorystore.Store
}

func NewUserRepository(usersDB inmemorystore.Store) repository.UserRepository {
	return &userRepository{usersDB: usersDB}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	now := time.Now()
	user.CreatedAt = &now
	r.usersDB.SaveWithID(user.ID, *user)
	return nil
}

func (r *userRepository) Get(ctx context.Context, userID int64) (*domain.User, error) {
	item, err := r.usersDB.Get(userID)
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
