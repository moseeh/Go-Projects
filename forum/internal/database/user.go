package database

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID         string
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
        uuid TEXT UNIQUE NOT NULL,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        fullname TEXT NOT NULL,
		session_id TEXT,
        session_expires_at DATETIME
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
	// Generate a new UUID for the user
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
    INSERT INTO users (uuid, username, email, password_hash, fullname)
    VALUES (?, ?, ?, ?, ?)`

	_, err = ur.DB.Exec(query, newUUID.String(), user.Username, user.Email, string(hashedPassword), user.Fullname)
	if err != nil {
		return err
	}

	user.UUID = newUUID.String()
	return nil
}

// getUserByUsername retrieves a user from the database by username
func (ur *UserRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := "SELECT uuid, username, email, password_hash, fullname FROM users WHERE username = ?"

	err := ur.DB.QueryRow(query, username).Scan(
		&user.UUID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Fullname,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) AddSession(username string) error {
	sessionID := generateSessionID()
	sessionExpiresAt := time.Now().Add(30 * time.Minute)
	_, err := ur.DB.Exec(
		"UPDATE users SET session_id = ?, session_expires_at = ? WHERE username = ?",
		sessionID, sessionExpiresAt, username,
	)
	return err
}

// generateSessionID creates a random session ID.
func generateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
