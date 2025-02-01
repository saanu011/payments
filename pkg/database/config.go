package database

import (
	"fmt"
	"strings"
)

type Config struct {
	Driver   string
	Name     string
	Host     string
	Port     int
	Username string
	Password string
	Query    string
}

func (cfg Config) ConnectionString() string {
	var kvs []string

	if len(cfg.Username) > 0 {
		kvs = append(kvs, fmt.Sprintf("user=%s", cfg.Username))
	}

	if len(cfg.Password) > 0 {
		kvs = append(kvs, fmt.Sprintf("password=%s", cfg.Password))
	}

	if len(cfg.Host) > 0 {
		kvs = append(kvs, fmt.Sprintf("host=%s", cfg.Host))
	}

	if cfg.Port > 0 {
		kvs = append(kvs, fmt.Sprintf("port=%d", cfg.Port))
	}

	if len(cfg.Name) > 0 {
		kvs = append(kvs, fmt.Sprintf("dbname=%s", cfg.Name))
	}

	if len(cfg.Query) > 0 {
		queries := strings.Split(cfg.Query, "&")
		kvs = append(kvs, queries...)
	}

	return strings.Join(kvs, " ")
}
