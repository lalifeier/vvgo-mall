package sha1

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	assert.Equal(t, Encode("123456"), "7c4a8d09ca3762af61e59520943dc26494f8941b")
}
