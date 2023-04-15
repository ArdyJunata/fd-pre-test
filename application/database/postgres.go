package database

import (
	"database/sql"
	"fd-test/application/config"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Postgres *sql.DB
}

func NewDB() DB {
	return DB{}
}

func (d DB) ConnectPostgres() DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetString(config.CFG_DB_HOST, "localhost"),
		config.GetString(config.CFG_DB_PORT, "5435"),
		config.GetString(config.CFG_DB_USER, "root"),
		config.GetString(config.CFG_DB_PASS, "root"),
		config.GetString(config.CFG_DB_NAME, "fd"),
		config.GetString(config.CFG_DB_SSL, "disable"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	d.Postgres = db

	return d
}
