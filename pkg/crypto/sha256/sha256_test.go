package sha256

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, Encode("123456"), "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92")

	assert.Equal(t, HmacEncode("123456", "test"), "8e46a2856fb67b0ce772fb6e7c5f31a8116adcdd4f1deafe4950e2cbb6609256")
}
