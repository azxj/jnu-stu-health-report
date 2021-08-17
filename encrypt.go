package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

func PKCS7Padding(text []byte, bs int) []byte {
	padding := bs - len(text)%bs
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

func Encrypt(plaintext string) (ciphertext string) {
	plainbytes := []byte(plaintext)
	key, iv := []byte("xAt9Ye&SouxCJziN"), []byte("xAt9Ye&SouxCJziN")
	block, _ := aes.NewCipher(key)
	plainbytes = PKCS7Padding(plainbytes, block.BlockSize())
	encrypter := cipher.NewCBCEncrypter(block, iv)
	cipherbytes := make([]byte, len(plainbytes))
	encrypter.CryptBlocks(cipherbytes, plainbytes)

	ciphertext = base64.StdEncoding.EncodeToString(cipherbytes)
	ciphertext = strings.ReplaceAll(ciphertext, "/", "_")
	ciphertext = strings.ReplaceAll(ciphertext, "+", "-")
	ciphertext = strings.Replace(ciphertext, "=", "*", 1)
	return ciphertext
}
