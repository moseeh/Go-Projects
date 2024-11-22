package handlers

import (
	"html/template"
	"net/http"

	"forum/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	UserRepo *database.UserRepository
}

type LoginCredentials struct {
	Username string
	Password string
}

func NewLoginHandler(userRepo *database.UserRepository) *Login {
	return &Login{
		UserRepo: userRepo,
	}
}

func (h *Login) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		}
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		user, err := h.UserRepo.GetUserByUsername(username)
		if err != nil {
			http.Error(w, "Error Finding User", http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "Invalid Username or password", http.StatusUnauthorized)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			// Password doesn't match
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
