package md5

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	assert.Equal(t, Encode("123456"), "e10adc3949ba59abbe56e057f20f883e")
}
