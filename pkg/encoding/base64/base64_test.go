package base64

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	str, _ := Decode("MTIzNDU2")
	assert.Equal(t, Encode("123456"), "MTIzNDU2")
	assert.Equal(t, str, "123456")
}
