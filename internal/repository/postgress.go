package repository

import (
	"database/sql"
	"fmt"
	"os"
)

type Repository struct {
	Mydb *sql.DB
}

func New() *Repository {
	return &Repository{}
}

func (db *Repository) Connect() error {
	user := os.Getenv("user")
	password := os.Getenv("password")
	host := os.Getenv("host")
	dbname := os.Getenv("dbname")
	sslmode := os.Getenv("sslmode")

	line := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=%s connect_timeout=5",
		host, dbname, user, password, sslmode)

	var err error
	db.Mydb, err = sql.Open("postgres", line)
	if err != nil {
		return err
	}

	if err = db.Mydb.Ping(); err != nil {
		return err
	}
	return nil
}

func (db *Repository) Close() error {
	if err := db.Mydb.Close(); err != nil {
		return err
	}
	return nil
}
