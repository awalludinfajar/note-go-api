package controller

import (
	"encoding/json"
	"net/http"

	"github.com/awalludinfajar/note-go-api.git/app/service"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequet struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:confirmPassword`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := service.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !user {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	service.CreateSession(w, req.Username)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequet
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if req.Password != req.ConfirmPassword {
		http.Error(w, "password doesnt match with confirm password", http.StatusBadGateway)
		return
	}

	user, err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, "password doesnt match with confirm password", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	service.DestroySession(w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}
