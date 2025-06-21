package domain

import (
	"time"
)

const (
	MaxPostLength = 240

	MaxLimit = int64(5)
)

type Tweet struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	Post      string     `json:"post"`
	CreatedAt *time.Time `json:"created_at"`
}

type TweetSearchParams struct {
	Limit   int64
	Offset  int64
	UserIDs []int64
}

type TweetsSearchResult struct {
	Paging  Paging  `json:"paging"`
	Results []Tweet `json:"results"`
}

type Paging struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
