package mangeFile

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type ManageFile struct {
	InputFilepath  string
	OutPutFilePath string
}

func (manageFiles *ManageFile) LoadDataFromFile() ([]string, error) {
	finalName := manageFiles.InputFilepath

	file, err := os.Open(finalName)
	if err != nil {

		return nil, errors.New("failed to load files")
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {

		return nil, errors.New("failed to read files")
	}

	return lines, nil
}

func (manageFiles *ManageFile) WriteJson(data interface{}) error {
	file, err := os.Create(manageFiles.OutPutFilePath)
	if err != nil {

		return errors.New("faild to create file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {

		return errors.New("faild to encode file")

	}

	time.Sleep(3 * time.Second)
	return nil
}

func NewManageFile(InputFilepath string, OutptFilepath string) ManageFile {
	return ManageFile{
		InputFilepath:  InputFilepath,
		OutPutFilePath: OutptFilepath,
	}
}
