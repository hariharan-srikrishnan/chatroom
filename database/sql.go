package database

import (
	"database/sql"
)

const (
	MYSQL = iota
)

func Connect(url string, database int) (*sql.DB, error) {
	// TODO: add support for mysql and sqlite3
	return nil, nil 
}
