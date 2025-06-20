package domain

import "time"

type User struct {
	ID        int64      `json:"id"`
	NickName  string     `json:"nick_name"`
	CreatedAt *time.Time `json:"created_at"`
}
