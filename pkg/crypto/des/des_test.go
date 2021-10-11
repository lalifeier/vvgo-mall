package des

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	key := []byte("Dlinkbook1qaz*WSX")

	encryptData, _ := Encrypt([]byte("123456"), key)
	decryptData, _ := Decrypt([]byte("b80a17bdec2dee79"), key)
	assert.Equal(t, string(encryptData), "b80a17bdec2dee79")
	assert.Equal(t, decryptData, "123456")
}
