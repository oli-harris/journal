package notes

import "time"

type NoteInternal struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	UserID    int64     `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Note struct {
	UUID      string    `json:"uuid"`
	UserID    int64     `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
