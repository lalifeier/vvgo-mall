package aes

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	var key = []byte("B31F2A75FBF94099")
	var iv = []byte("1234567890123456")

	encryptData, _ := Encrypt([]byte("123456"), key, iv)
	decryptData, _ := Decrypt("NBIiLNmTcLO7gn3SapDtrQ==", key, iv)
	assert.Equal(t, encryptData, "NBIiLNmTcLO7gn3SapDtrQ==")
	assert.Equal(t, decryptData, "123456")
}
