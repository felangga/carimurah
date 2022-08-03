package engines

import (
	"carimurah/postgres"
)

type Engine struct {
	db *postgres.Postgres
}

func NewEngine(db *postgres.Postgres) *Engine {
	return &Engine{
		db: db,
	}
}
