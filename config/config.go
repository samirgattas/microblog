package config

import (
	"os"

	"github.com/samirgattas/microblog/internal/core/domain"
	inmemorystore "github.com/samirgattas/microblog/lib/in_memory_store"
)

var c *Config

type Config struct {
	UsersDB    inmemorystore.Store
	FollowedDB map[int64]domain.Followed
	TweetDB    map[int64]domain.Tweet
	Domain     string
	PortDB     int64
	DomainDB   string
	UserDB     string
	PasswordDB string
}

func (c *Config) NewConfig(userDB inmemorystore.Store, followedDB map[int64]domain.Followed, tweetDB map[int64]domain.Tweet) *Config {
	domain := "localhost"
	if isDockerEnv() {
		domain = ""
	}
	c = &Config{
		UsersDB:    userDB,
		FollowedDB: followedDB,
		TweetDB:    tweetDB,
		Domain:     domain,
		PortDB:     3307,
		DomainDB:   "127.0.0.1",
		UserDB:     "root",
		PasswordDB: "root",
	}
	return c
}

func isDockerEnv() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
