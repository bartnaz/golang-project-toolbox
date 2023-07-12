package jsonprocessor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ProcessJSON[T any](filepath string, body *T) *T {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, body)
	return body
}
