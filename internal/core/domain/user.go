package domain

import "time"

type User struct {
	ID        int64      `json:"id"`
	Nickname  string     `json:"nickname"`
	CreatedAt *time.Time `json:"created_at"`
}
