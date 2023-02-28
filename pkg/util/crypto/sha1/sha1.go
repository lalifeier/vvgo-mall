package sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
