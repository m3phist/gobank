// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"
)

type User struct {
	ID            int64     `json:"id"`
	Email         string    `json:"email"`
	HashedPasswrd string    `json:"hashed_passwrd"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
