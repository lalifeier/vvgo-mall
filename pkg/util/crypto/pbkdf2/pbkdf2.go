package pbkdf2

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 4096, sha256.Size, sha256.New)
}

//func HashPassword(str string) string {
//	hashedPassword := pbkdf2.Key(stringToUTF16Bytes(str), stringToUTF16Bytes(PasswordSecuritySalt), PasswordSecurityIterations, PasswordSecurityKeylen, sha1.New)
//	return b64.StdEncoding.EncodeToString(hashedPassword)
//}
