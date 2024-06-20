package models

import (
	"time"
)

type User struct {
    UserID    string    `db:"user_id" json:"user_id"`
    Name      string    `db:"name" json:"name"`
    Email     string    `db:"email" json:"email"`
    Birthday  time.Time `db:"birthday" json:"birthday"`
    Password  string    `db:"password" json:"password"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
    DeletedAt int64     `db:"deleted_at" json:"deleted_at"`
}