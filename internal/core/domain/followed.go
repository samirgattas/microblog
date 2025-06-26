package domain

import "time"

type Followed struct {
	ID             int64      `json:"id"`
	UserID         int64      `json:"user_id"`
	FollowedUserID int64      `json:"followed_user_id"`
	Enabled        bool       `json:"enabled"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type FollowedPatchCommand struct {
	Enabled *bool `json:"enabled"`
}
