package service

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/awalludinfajar/note-go-api.git/app/model"
	"github.com/google/uuid"
)

const noteFile = "./data/notes.json"

func loadNote() ([]model.Note, error) {
	if _, err := os.Stat(noteFile); os.IsNotExist(err) {
		return []model.Note{}, nil
	}

	data, err := os.ReadFile(noteFile)

	if err != nil {
		return nil, err
	}

	var notes []model.Note
	if len(data) > 0 {
		err = json.Unmarshal(data, &notes)
		if err != nil {
			return nil, err
		}
	}
	return notes, nil
}

func saveNotes(notes []model.Note) error {
	data, err := json.Marshal(notes)
	if err != nil {
		return err
	}
	return os.WriteFile(noteFile, data, 0644)
}

func GetAllNote() ([]model.Note, error) {
	return loadNote()
}

func CreateNote(title, content string) (model.Note, error) {
	notes, err := loadNote()
	if err != nil {
		return model.Note{}, err
	}

	newNote := model.Note{
		ID:      int(uuid.Must(uuid.NewRandom())[0]),
		Title:   title,
		Content: content,
	}

	notes = append(notes, newNote)
	if err := saveNotes(notes); err != nil {
		return model.Note{}, err
	}

	return newNote, nil
}

func UpdateNote(Id, title, content string) (model.Note, error) {
	notes, err := loadNote()
	if err != nil {
		return model.Note{}, err
	}

	for i, note := range notes {
		if strconv.Itoa(note.ID) == Id {
			notes[i].Title = title
			notes[i].Content = content
			if err := saveNotes(notes); err != nil {
				return model.Note{}, err
			}
			return notes[i], nil
		}
	}

	return model.Note{}, errors.New("note not found")
}

func DeleteNote(id string) error {
	notes, err := loadNote()
	if err != nil {
		return err
	}

	for i, note := range notes {
		if strconv.Itoa(note.ID) == id {
			notes = append(notes[:i], notes[i+1:]...)
			return saveNotes(notes)
		}
	}

	return errors.New("note not found")
}
