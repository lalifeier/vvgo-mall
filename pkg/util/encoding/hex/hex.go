package hex

import (
	"encoding/hex"
)

func Encode(str string) string {
	return hex.EncodeToString([]byte(str))
}

func Decode(str string) (string, bool) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}
