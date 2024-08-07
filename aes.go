package svc

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(plainText []byte) (string, error) {
	key := []byte(AesKey())
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, len(plainText))
	stream := cipher.NewCTR(block, key)
	stream.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText[:]), nil
}

func AesDecrypt(cipherBytes []byte) ([]byte, error) {
	decodeText, err := base64.StdEncoding.DecodeString(string(cipherBytes))
	key := []byte(AesKey())
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(decodeText))
	stream := cipher.NewCTR(block, key)
	stream.XORKeyStream(plainText, decodeText[:])
	return plainText, nil
}
