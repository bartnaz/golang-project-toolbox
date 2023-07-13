package csvprocessor

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

func ProcessCSV[T any](filepath string, body []*T) []*T {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &body); err != nil {
		fmt.Println(err)
	}

	return body
}
