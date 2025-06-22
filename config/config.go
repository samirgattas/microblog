package config

import inmemorystorage "github.com/samirgattas/microblog/internal/core/lib/customerror/in_memory_storage"

var c *Config

type Config struct {
	UserDB inmemorystorage.Store
}

func (c *Config) NewConfig(userDB inmemorystorage.Store) *Config {
	if c != nil {
		return c
	}
	c = &Config{
		UserDB: userDB,
	}
	return c
}
