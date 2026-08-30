package notification

import (
	"Service/internal/models"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

const StoreSQLCommand = `INSERT INTO Logs(token, chat_id, message, ip) VALUES ($1, $2, $3, $4)`

type Repository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) StoreData(ctx context.Context, req models.WebhookReceiver) error {
	_, err := r.db.ExecContext(ctx, StoreSQLCommand, req.Token, req.ChatID, req.Message, req.IP)
	if err != nil {
		return err
	}
	return nil
}
