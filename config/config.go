package config

import (
	"github.com/samirgattas/microblog/internal/core/domain"
	inmemorystore "github.com/samirgattas/microblog/lib/in_memory_store"
)

var c *Config

type Config struct {
	UserDB     inmemorystore.Store
	FollowedDB map[int64]domain.Followed
	TweetDB    map[int64]domain.Tweet
}

func (c *Config) NewConfig(userDB inmemorystore.Store, followedDB map[int64]domain.Followed, tweetDB map[int64]domain.Tweet) *Config {
	c = &Config{
		UserDB:     userDB,
		FollowedDB: followedDB,
		TweetDB:    tweetDB,
	}
	return c
}
