package crc32

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	assert.Equal(t, Encode("123456"), uint32(158520161))
}
