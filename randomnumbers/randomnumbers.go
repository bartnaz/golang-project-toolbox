package randomnumbers

import (
	"encoding/base64"
	"math/rand"
	"strconv"
)

func GenerateStringNumber(l int) (s string) {
	for i := 0; i < l; i++ {
		s += strconv.Itoa(rand.Intn(9))[:1]
	}
	return
}

func EncodeValues(values ...string) (s string) {
	for _, v := range values {
		s += v
	}
	return base64.StdEncoding.EncodeToString([]byte(s))
}
