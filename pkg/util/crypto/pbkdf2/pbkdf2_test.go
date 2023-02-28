package pbkdf2

import (
	"testing"
)

func Test(t *testing.T) {
	pass := []byte("foo")
	salt := []byte("bar")

	t.Log(HashPassword(pass, salt))
}
