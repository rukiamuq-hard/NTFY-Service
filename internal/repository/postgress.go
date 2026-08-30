package repository

import (
	"database/sql"
	"fmt"
	"os"
)

type Postgress struct {
	DB *sql.DB
}

func New() *Postgress {
	return &Postgress{}
}

func (pg *Postgress) Connect() error {
	user := os.Getenv("user")
	password := os.Getenv("password")
	host := os.Getenv("host")
	dbname := os.Getenv("dbname")
	sslmode := os.Getenv("sslmode")

	line := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=%s connect_timeout=5",
		host, dbname, user, password, sslmode)

	var err error
	pg.DB, err = sql.Open("postgres", line)
	if err != nil {
		return err
	}

	if err = pg.DB.Ping(); err != nil {
		return err
	}
	return nil
}

func (pg *Postgress) Close() error {
	if err := pg.DB.Close(); err != nil {
		return err
	}
	return nil
}
