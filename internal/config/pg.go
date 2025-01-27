package config

import (
	"errors"
	"os"
)

type PGConfig interface {
	GetDsn() string
}

type pg struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, errors.New("DB_DSN environment variable not set")
	}

	return &pg{dsn: dsn}, nil
}

func (pg *pg) GetDsn() string {
	return pg.dsn
}
