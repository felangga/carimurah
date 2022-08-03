package postgres

import (
	"carimurah/configs"

	"database/sql"
	"fmt"
)

type Postgres struct {
	DB *sql.DB
}

func New() (*Postgres, error) {
	db, err := sql.Open("postgres", configs.GetDBConfig())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database: Successfully connected!")
	return &Postgres{
		DB: db,
	}, nil
}
