// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt sql.NullTime
}