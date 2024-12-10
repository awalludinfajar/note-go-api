package service

import (
	"errors"
	"net/http"
)

var validUsername = "admin"
var validPassword = "password"

func AuthenticateUser(username, password string) (bool, error) {
	users, err := loadUser()
	if err != nil {
		return false, errors.New("unable to load users: " + err.Error())
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true, nil
		}
	}
	return false, nil
}

func CreateSession(w http.ResponseWriter, username string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: username,
		Path:  "/",
	})
}

func ValidateSession(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", errors.New("no valid session")
	}
	return cookie.Value, nil
}

func DestroySession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
