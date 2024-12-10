package service

import (
	"encoding/json"
	"os"

	"github.com/awalludinfajar/note-go-api.git/app/model"
	"github.com/google/uuid"
)

const userData = "./data/user.json"

func loadUser() ([]model.User, error) {
	if _, err := os.Stat(noteFile); os.IsNotExist(err) {
		return []model.User{}, nil
	}

	data, err := os.ReadFile(noteFile)

	if err != nil {
		return nil, err
	}

	var user []model.User
	if len(data) > 0 {
		err = json.Unmarshal(data, &user)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func saveUser(user []model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return os.WriteFile(noteFile, data, 0644)
}

func RegisterUser(username, password string) (model.User, error) {
	user, err := loadUser()
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		ID:       int(uuid.Must(uuid.NewRandom())[0]),
		Username: username,
		Password: password,
	}

	user = append(user, newUser)
	if err := saveUser(user); err != nil {
		return model.User{}, err
	}

	return newUser, nil
}
