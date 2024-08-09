package randomnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateStringNumber(t *testing.T) {

	result := GenerateStringNumber(7)
	assert.Equal(t, 7, len(result))
}

func TestEncodeValues(t *testing.T) {

	result := EncodeValues("test string")
	assert.NotNil(t, len(result))
}
