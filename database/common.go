package database

import (
	"database/sql"
)

type Connector struct {
	conn *sql.DB	
}

func (c *Connector) Query(query string, opts ...interface{}) (*sql.Rows, error) {
	rows, err := c.conn.Query(query, opts)
	return rows, err
}
