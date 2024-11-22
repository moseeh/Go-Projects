package main

import (
	"database/sql"
	"log"
	"net/http"

	"forum/internal/database"
	"forum/internal/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Initialize repositories
	userRepo := database.NewUserRepository(db)
	err = userRepo.CreateUsersTable()
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}
	signupHandler := handlers.NewSignupHandler(userRepo)
	loginHandler := handlers.NewLoginHandler(userRepo)

	// Initialize handler
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", loginHandler.LoginHandler)
	http.HandleFunc("/signup", signupHandler.SignupHandler)
	log.Println("The server is running on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
