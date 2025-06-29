package followed

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
	insertFollowedQuery  = `INSERT INTO Followed (user_id, followed_user_id, enabled, created_at, updated_at) VALUES (?,?,?,?,?);`
	getFollowedByIDQuery = `SELECT id, user_id, followed_user_id, enabled, created_at, updated_at FROM Followed WHERE id = ?;`
	updateFollowedQuery  = `UPDATE Followed SET enabled = ?, updated_at = ? WHERE id = ?;`
)

type followedRepository struct {
	followedDB *sql.DB
}

func NewFollowedRepository(followedDB *sql.DB) repository.FollowedRepository {
	return &followedRepository{
		followedDB: followedDB,
	}
}

func (r *followedRepository) Save(ctx context.Context, followed *domain.Followed) error {
	now := time.Now()
	followed.CreatedAt = &now
	followed.UpdatedAt = &now
	res, err := r.followedDB.ExecContext(ctx, insertFollowedQuery, followed.UserID, followed.FollowedUserID, followed.Enabled, followed.CreatedAt, followed.UpdatedAt)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// Update item id
	followed.ID = lastID
	return nil
}

func (r *followedRepository) Get(ctx context.Context, ID int64) (*domain.Followed, error) {
	row := r.followedDB.QueryRowContext(ctx, getFollowedByIDQuery, ID)

	followed := domain.Followed{}
	var createdAt, updatedAt string
	err := row.Scan(&followed.ID, &followed.UserID, &followed.FollowedUserID, &followed.Enabled, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.Followed{}, customerror.NewNotFoundError("followed")
		}
		return &domain.Followed{}, customerror.NewInternalServerError(err.Error())
	}

	tmpCreatedAt, err := time.Parse(time.DateTime, createdAt)
	if err != nil {
		return &domain.Followed{}, customerror.NewInternalServerError(err.Error())
	}

	tmpUpdatedAt, err := time.Parse(time.DateTime, updatedAt)
	if err != nil {
		return &domain.Followed{}, customerror.NewInternalServerError(err.Error())
	}

	followed.CreatedAt = &tmpCreatedAt
	followed.UpdatedAt = &tmpUpdatedAt

	return &followed, nil
}

func (r *followedRepository) Update(ctx context.Context, followed *domain.Followed) error {
	now := time.Now()
	followed.UpdatedAt = &now
	res, err := r.followedDB.ExecContext(ctx, updateFollowedQuery, followed.Enabled, followed.UpdatedAt, followed.ID)
	if err != nil {
		return err
	}

	number, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if number == 0 {
		return customerror.NewNotFoundError("followed")
	}

	return nil
}

func (r *followedRepository) SearchByUserIDAndFollowedUserID(ctx context.Context, followerUserID *int64, followedUserID *int64) ([]domain.Followed, error) {
	searchQuery := `SELECT id, user_id, followed_user_id, enabled, created_at, updated_at FROM Followed`

	var args []interface{}
	if followerUserID != nil || followedUserID != nil {
		searchQuery += ` WHERE `
		if followerUserID != nil {
			searchQuery += `user_id = ?`
			args = append(args, *followerUserID)
		}
		if followerUserID != nil && followedUserID != nil {
			searchQuery += ` AND `
		}
		if followedUserID != nil {
			searchQuery += `followed_user_id = ?`
			args = append(args, *followedUserID)
		}

	}

	followedArray := []domain.Followed{}
	rows, err := r.followedDB.QueryContext(ctx, searchQuery, args...)
	if err != nil {
		return []domain.Followed{}, err
	}
	for rows.Next() {
		followed := domain.Followed{}
		var createdAt, updatedAt string
		err := rows.Scan(&followed.ID, &followed.UserID, &followed.FollowedUserID, &followed.Enabled, &createdAt, &updatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return []domain.Followed{}, customerror.NewNotFoundError("followed")
			}
			return []domain.Followed{}, customerror.NewInternalServerError(err.Error())
		}

		tmpCreatedAt, err := time.Parse(time.DateTime, createdAt)
		if err != nil {
			return []domain.Followed{}, customerror.NewInternalServerError(err.Error())
		}

		tmpUpdatedAt, err := time.Parse(time.DateTime, updatedAt)
		if err != nil {
			return []domain.Followed{}, customerror.NewInternalServerError(err.Error())
		}

		followed.CreatedAt = &tmpCreatedAt
		followed.UpdatedAt = &tmpUpdatedAt

		followedArray = append(followedArray, followed)
	}
	return followedArray, nil
}
