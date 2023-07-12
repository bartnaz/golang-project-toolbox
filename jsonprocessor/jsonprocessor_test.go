package jsonprocessor

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

type TestStruct struct {
	ID   int    `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

func TestProcessJSON(t *testing.T) {
	filePath, _ := filepath.Abs("test_file.json")

	body := TestStruct{}
	returnValue := ProcessJSON[TestStruct](filePath, &body)

	assert.NotNil(t, returnValue)
	assert.Equal(t, 1, returnValue.ID)
	assert.Equal(t, "Test", returnValue.Name)
}
