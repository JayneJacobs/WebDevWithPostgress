package model

import (
	"database/sql"
)

var db DB

// DB Methods for test SQL package relys on concrete types
// need an interface that allows us to inject a tesable mock
type DB interface {
	QueryRow(string, ...interface{}) Row
	Exec(string, ...interface{}) (Result, error)
}

// SetDatabase points to the connection in Main from the Model Layer

//Row local method
type Row interface {
	Scan(...interface{}) error
}

// Result is the Id of last login record
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type sqlDB struct {
	db *sql.DB
}

func (s sqlDB) QueryRow(query string, args ...interface{}) Row {
	return s.db.QueryRow(query, args...)
}
func (s sqlDB) Exec(query string, args ...interface{}) (Result, error) {
	return s.db.Exec(query, args...)
}

//SetDatabase method takes database name and pointer to DB
func SetDatabase(database *sql.DB) {
	db = &sqlDB{database}
}
