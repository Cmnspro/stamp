package local

import (
	"database/sql"
)

type Service struct {
	db *sql.DB
}

func NewService(exec *sql.DB) *Service {
	return &Service{
		db: exec,
	}
}
