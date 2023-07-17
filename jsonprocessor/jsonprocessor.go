package jsonprocessor

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ProcessJSON[T any](filepath string, body *T) (*T, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
