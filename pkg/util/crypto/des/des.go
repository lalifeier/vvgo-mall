package des

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockModel := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	ciphertext := make([]byte, len(origData))
	blockModel.CryptBlocks(ciphertext, origData)
	return ciphertext, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Decrypt(ciphertext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockModel := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(origData, ciphertext)
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return "", err
	}
	return string(origData), nil
}

func PKCS7UnPadding(plantText []byte) ([]byte, error) {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)], nil
}
