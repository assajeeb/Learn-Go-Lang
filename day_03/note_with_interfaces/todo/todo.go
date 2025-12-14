package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("no data inputed")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println("text: ", todo.Text)

}

func (todo Todo) Save() error {
	if todo.Text == "" {
		return errors.New("title and content not found for saving")
	}
	err := todo.SaveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (todo Todo) SaveToFile() error {
	filename := strings.TrimSuffix(todo.Text, "\n")
	filename = strings.ReplaceAll(filename, "\n", " ")
	filename = strings.ReplaceAll(filename, "\r", " ")
	filename = filename + ".json"

	jsonData, err := json.Marshal(todo)
	if err != nil {
		return errors.New("faild in parsing the note")
	}

	err = os.WriteFile(filename, jsonData, 0644)

	if err != nil {
		return err
	}
	return nil
}
