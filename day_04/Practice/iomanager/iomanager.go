package iomanager

type IoManger interface {
	LoadDataFromFile() ([]string, error)
	WriteJson(data interface{}) error
}
