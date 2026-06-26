package postgres

import (
	"context"
	"service-wedding/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type contactRepo struct {
	pool *pgxpool.Pool
}

func NewContactRepository(pool *pgxpool.Pool) domain.ContactRepository {
	return &contactRepo{pool: pool}
}

func (r *contactRepo) Create(ctx context.Context, msg *domain.ContactMessage) error {
	query := `INSERT INTO contact_messages (name, email, message, created_at) 
	          VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id, created_at`
	err := r.pool.QueryRow(ctx, query, msg.Name, msg.Email, msg.Message).Scan(&msg.ID, &msg.CreatedAt)
	return err
}

func (r *contactRepo) GetAll(ctx context.Context) ([]domain.ContactMessage, error) {
	query := `SELECT id, name, email, message, created_at FROM contact_messages ORDER BY id DESC`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []domain.ContactMessage
	for rows.Next() {
		var msg domain.ContactMessage
		err := rows.Scan(&msg.ID, &msg.Name, &msg.Email, &msg.Message, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
