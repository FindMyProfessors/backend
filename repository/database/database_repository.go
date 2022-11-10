package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DatabasePool *pgxpool.Pool
}
