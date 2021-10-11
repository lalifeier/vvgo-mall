package jwtauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	jwtAuth := New()

	// userID := 123
	userName := "name"
	token, err := jwtAuth.GenerateToken(1, userName)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
