package user

import (
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

func (u *userRepository) Save(user *domain.User) error {
	now := time.Now()
	user.CreatedAt = &now
	u.usersDB[user.ID] = *user
	return nil
}

func (u *userRepository) Get(ID int64) (*domain.User, error) {
	user := domain.User{}
	user = u.usersDB[ID]
	if user.ID == 0 {
		slog.Error("user not found")
		return &domain.User{}, customerror.NewNotFoundError("user")
	}
	return &user, nil
}
