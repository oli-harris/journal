package db

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:"id,autoincrement,pk"`
	UUID      string    `bun:"uuid,unique,notnull"`
	Email     string    `bun:"email,unique,notnull"`
	Password  string    `bun:"password,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Note struct {
	ID        int64     `bun:"id,autoincrement,pk"`
	UUID      string    `bun:"uuid,unique,notnull"`
	UserID    int64     `bun:"user_id,notnull"`
	User      *User     `bun:"rel:belongs-to,join:user_id=id"`
	Note      string    `bun:"note,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
