package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server   `yaml:"server"`
	MySQL  Database `yaml:"mysql"`
}

type Server struct {
	Domain string `yaml:"domain"`
	Port   int64  `yaml:"port"`
}

type Database struct {
	Domain   string `yaml:"domain"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (c *Config) NewConfig() *Config {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	data, err := ioutil.ReadFile(pwd + "/../config/env_local.yaml")
	if err != nil {
		log.Fatalf("Cannot read config file")
	}

	var cfg *Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Cannot unmarshal config")
	}

	domain := "localhost"
	if isDockerEnv() {
		domain = ""
	}
	cfg.Server.Domain = domain
	return cfg
}

func isDockerEnv() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
