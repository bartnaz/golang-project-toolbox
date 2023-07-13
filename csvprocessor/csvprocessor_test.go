package csvprocessor

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

type TestStruct struct {
	AdvCID  int    `csv:"Adv_CID"`
	PubName string `csv:"Pub Name"`
}

func TestProcessJSON(t *testing.T) {
	filePath, _ := filepath.Abs("test_file.csv")

	body := []*TestStruct{}
	returnValue := ProcessCSV[TestStruct](filePath, body)

	assert.NotNil(t, returnValue)
	assert.Equal(t, 6111112, returnValue[0].AdvCID)
	assert.Equal(t, "TestTECH PTE. LTD.", returnValue[0].PubName)

}
