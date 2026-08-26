package user

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

const StoreSQLCommand = `INSERT INTO Logs(token, chat_id, message) VALUES ($1, $2, $3)`

type Repository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) StoreData(ctx context.Context, token string, chat_id int64, message string) error {
	_, err := r.db.ExecContext(ctx, StoreSQLCommand, token, chat_id, message)
	if err != nil {
		return err
	}
	return nil
}
