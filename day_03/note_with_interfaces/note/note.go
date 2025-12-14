package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Contant   string    `json:"content"`
	CreatedAt time.Time `json:"create_at"`
}

func New(title, contant string) (Note, error) {
	if title == "" || contant == "" {
		return Note{}, errors.New("no data inputed")
	}
	return Note{
		Title:     title,
		Contant:   contant,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Contant)
}

func (note Note) Save() error {
	if note.Title == "" || note.Contant == "" {
		return errors.New("title and content not found for saving")
	}
	err := note.SaveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (note Note) SaveToFile() error {
	filename := strings.TrimSuffix(note.Title, "\n")
	filename = strings.ReplaceAll(filename, "\n", " ")
	filename = strings.ReplaceAll(filename, "\r", " ")
	filename = filename + ".json"

	jsonData, err := json.Marshal(note)
	if err != nil {
		return errors.New("faild in parsing the note")
	}

	os.WriteFile(filename, jsonData, 0644)
	return nil
}
