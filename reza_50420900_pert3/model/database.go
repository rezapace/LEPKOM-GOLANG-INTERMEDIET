package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Table interface {
	Name() string
	Field() ([]string, []interface{})
}

func Connect(username string, password string, host string, database string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, database)
	db, err := sql.Open("mysql", conn)
	return db, err
}

func CreateDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("CREATE DATABASE %v", name)
	_, err := db.Exec(query)
	return err
}

func CreateTable(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	return err
}

func DropDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("DROP DATABASE  %v", name)
	_, err := db.Exec(query)
	return err
}
