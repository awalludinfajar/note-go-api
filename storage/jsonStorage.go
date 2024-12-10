package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/awalludinfajar/note-go-api.git/model"
)

type JSONStorage struct {
	Filepath string
	lock     sync.Mutex
}

func newJosnStorage(filePath string) *JSONStorage {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.WriteFile(filePath, []byte("[]"), 0644)
	}

	return &JSONStorage{Filepath: filePath}
}

func (s *JSONStorage) LoadNote() ([]model.Note, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	file, err := os.Open(s.Filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var notes []model.Note
	if err := json.NewDecoder(file).Decode(&notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (s *JSONStorage) SaveNotes(notes []model.Note) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	file, err := os.OpenFile(s.Filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(notes); err != nil {
		return err
	}
	return nil
}
