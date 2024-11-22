package database

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	Fullname     string
}

type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates the users table if it doesn't exist
func (ur *UserRepository) CreateUsersTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        fullname TEXT NOT NULL
    )`

	_, err := ur.DB.Exec(query)
	return err
}

func (ur *UserRepository) UserExists(username, email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ? OR email = ?)"
	err := ur.DB.QueryRow(query, username, email).Scan(&exists)
	return exists, err
}

func (ur *UserRepository) CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
    INSERT INTO users (username, email, password_hash, fullname)
    VALUES (?, ?, ?, ?)`

	result, err := ur.DB.Exec(query, user.Username, user.Email, string(hashedPassword), user.Fullname)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}
