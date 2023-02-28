package hex

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	str, _ := Decode("616263")
	assert.Equal(t, Encode("abc"), "616263")
	assert.Equal(t, str, "abc")
}
