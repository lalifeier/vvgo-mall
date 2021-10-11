package rsa

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/lalifeier/vgo/pkg/encoding/base64"
)

//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = `
-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgHVJTb438xI/Dcr5XE5HQaTfL0Zj
KsFTxhpt9sC8ri21ZzEYXwkVw6IV2vE+NGW34YU8ogqIXxav6tji94dKjbuG3X9V
7LV6hssJkzQ2bjhhGsg1Vqzr9EPwwSuu7lAb6z4KQlBX6I2sRFCIBoMSXd+5GZVx
DgRuefNJ7tVl11OZAgMBAAE=
-----END PUBLIC KEY-----
`

//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgHVJTb438xI/Dcr5XE5HQaTfL0ZjKsFTxhpt9sC8ri21ZzEYXwkV
w6IV2vE+NGW34YU8ogqIXxav6tji94dKjbuG3X9V7LV6hssJkzQ2bjhhGsg1Vqzr
9EPwwSuu7lAb6z4KQlBX6I2sRFCIBoMSXd+5GZVxDgRuefNJ7tVl11OZAgMBAAEC
gYAF3VV61ndcCTUaWh+odl2s7dACWAESS6sTNT5qYsTe2jw/szVFUgZHO4qIF5Et
KOWo8OA1YJ5IlyFKGQdBh5PQpB/GpGJCZOmxv9nlZuRYz884Ari0aGlQBzvmbDvg
qibsO77YfR+A9+kPiD3OpusvZT5DIAbMHA/X3npDE4mRpQJBAMtOYLwVu5WfaoVF
kKf7GsBxrMWJuD1cqa7njVCLlde7H154kBEzOHLzfrXGyshME8xkdfAgilKJBS5O
vamNU4cCQQCTr2bFhEOgTq2oMCr0kosP0sJiYtvAqOkxr03k9p38CX+Q6Jcb+iZR
w9ssD4csUYSlQRXzTT3jssB17A94biffAkBxkFZ4uasO6P2XdERZkOpglR7tOQCx
RhGCodVOKKqK1vuuyamv8eyWSW1+HI0pVVW51mQKviKF+APs2g8XptoBAkEAhxYH
oV+sI/QTsCXvBKsOfDjCCRB4Ba/7LbE2RNq4A5QElV3K6pJTfrLxxUmm0Qj3ldkE
5PmgRKh6luKH/BZ2bQJAAiumrvjSLUYqhwZPPOJEKtP3ddh1nmwJNHvZzNvioW/n
fRKQwWZ0VbVA6WYApN9hEu7fKEBIu/7iRXc2mTEZDA==
-----END RSA PRIVATE KEY-----
`

func Test(t *testing.T) {
	var str = "123456"
	encryptBytes, _ := Encrypt([]byte(str), publicKey)
	decryptBytes, _ := Decrypt(encryptBytes, privateKey)
	t.Log(base64.Encode(string(encryptBytes)))
	assert.Equal(t, string(decryptBytes), str)

	signBytes, _ := Sign([]byte(str), privateKey)
	err := VerifySign([]byte(str), signBytes, publicKey)
	if err != nil {
		t.Log(err)
	}
}
