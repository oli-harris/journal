package notes

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	List(ctx context.Context, userID int64) ([]Note, error)
	Create(ctx context.Context, userID int64, note string) (string, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) Create(ctx context.Context, userID int64, note string) (string, error) {
	newUuid := uuid.NewString()

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO notes (uuid, user_id, note, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, newUuid, userID, note, time.Now().UTC(), time.Now().UTC())
	if err != nil {
		return "", err
	}

	return newUuid, nil
}

func (r *SQLRepository) List(ctx context.Context, userID int64) ([]Note, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT uuid, user_id, note, created_at, updated_at
		FROM notes
		WHERE user_id = ?
		ORDER BY created_at DESC
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.UUID, &n.UserID, &n.Note, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
