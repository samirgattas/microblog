package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/repository"
	"github.com/samirgattas/microblog/lib/customerror"
)

var (
	insertUserQuery  = `INSERT INTO User (user_id, nickname, created_at) VALUES (?,?,?);`
	getUserByIDQuery = `SELECT user_id, nickname, created_at FROM User WHERE user_id = ?;`
)

type userRepository struct {
	usersDB *sql.DB
}

func NewUserRepository(usersDB *sql.DB) repository.UserRepository {
	return &userRepository{usersDB: usersDB}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	now := time.Now()
	user.CreatedAt = &now
	_, err := r.usersDB.ExecContext(ctx, insertUserQuery, user.ID, user.Nickname, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Get(ctx context.Context, userID int64) (*domain.User, error) {
	user := domain.User{}
	row := r.usersDB.QueryRow(getUserByIDQuery, userID)
	var createdAt string
	err := row.Scan(&user.ID, &user.Nickname, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, customerror.NewNotFoundError("user")
		}
		return &domain.User{}, customerror.NewInternalServerError(err.Error())
	}

	tmpCreatedAt, err := time.Parse(time.DateTime, createdAt)
	if err != nil {
		return &domain.User{}, customerror.NewInternalServerError(err.Error())
	}
	user.CreatedAt = &tmpCreatedAt

	return &user, nil
}
