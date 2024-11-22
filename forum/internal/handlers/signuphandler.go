package handlers

import (
	"html/template"
	"net/http"

	"forum/internal/database"
)

type Signup struct {
	UserRepo *database.UserRepository
}

func NewSignupHandler(userRepo *database.UserRepository) *Signup {
	return &Signup{
		UserRepo: userRepo,
	}
}

func (h *Signup) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/signup.html")
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
		email := r.FormValue("email")
		password := r.FormValue("password")
		fullname := r.FormValue("fullname")
		if username == "" || email == "" || password == "" || fullname == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}
		exists, err := h.UserRepo.UserExists(username, email)
		if err != nil {
			http.Error(w, "Error checking user existence: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Username or email already exists", http.StatusConflict)
			return
		}
		user := &database.User{
			Username:     username,
			Email:        email,
			PasswordHash: password,
			Fullname:     fullname,
		}
		err = h.UserRepo.CreateUser(user)
		if err != nil {
			http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
