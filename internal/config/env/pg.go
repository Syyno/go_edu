package envconfig

import (
	"errors"
	"fmt"
)

const (
	dbName   = "POSTGRES_DATABASE_NAME"
	userName = "POSTGRES_USER"
	password = "POSTGRES_PASSWORD"
	port     = "POSTGRES_PORT"
	host     = "POSTGRES_HOST"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig(conf *Configuration) (PGConfig, error) {
	db, err := conf.Get(dbName)
	if err != nil {
		return nil, errors.New("pg db name not found")
	}

	user, err := conf.Get(userName)
	if err != nil {
		return nil, errors.New("pg user name not found")
	}

	pw, err := conf.Get(password)
	if err != nil {
		return nil, errors.New("pg password not found")
	}

	dbPort, err := conf.Get(port)
	if err != nil {
		return nil, errors.New("pg port not found")
	}

	hst, err := conf.Get(host)
	if err != nil {
		return nil, errors.New("pg host not found")
	}

	dsn := createDSN(db, user, pw, dbPort, hst)

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}

func createDSN(db, user, password, port, host string) string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host, port, db, user, password)
}
