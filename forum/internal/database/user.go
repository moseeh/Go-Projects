package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository struct {
	DB *sql.DB
}
