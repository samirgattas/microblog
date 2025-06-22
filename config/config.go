package config

import inmemorystore "github.com/samirgattas/microblog/internal/core/lib/customerror/in_memory_store"

var c *Config

type Config struct {
	UserDB inmemorystore.Store
}

func (c *Config) NewConfig(userDB inmemorystore.Store) *Config {
	if c != nil {
		return c
	}
	c = &Config{
		UserDB: userDB,
	}
	return c
}
