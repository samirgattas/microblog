package domain

import (
	"time"
)

var (
	MaxLimit = int64(5)
)

type Tweet struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	Post      string     `json:"post"`
	CreatedAt *time.Time `json:"created_at"`
}
