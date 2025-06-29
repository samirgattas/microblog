package config

import (
	"os"
)

var c *Config

type Config struct {
	MySQL  Database `yaml:"mysql"`
	Domain string   `yaml:"domain"`
}

type Database struct {
	Domain   string `yaml:"domain"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (c *Config) NewConfig() *Config {
	domain := "localhost"
	if isDockerEnv() {
		domain = ""
	}
	c = &Config{
		Domain: domain,
		MySQL: Database{
			Port:     3307,
			Domain:   "127.0.0.1",
			User:     "root",
			Password: "root",
		},
	}
	return c
}

func isDockerEnv() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
