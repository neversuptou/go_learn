package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/fatih/color"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	keyFromEnv := os.Getenv("KEY")
	if keyFromEnv == "" {
		color.Red("Внимание: используется ключ шифрования по умолчанию. Для повышения безопасности установите свой ключ в переменной окружения KEY")
		keyFromEnv = "4a8af831086259267ba5e2fd4cd328f5"
	}
	return &Encrypter{
		Key: keyFromEnv,
	}
}

func (enc *Encrypter) Encrypt(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainString, nil)
}

func (enc *Encrypter) Decrypt(encryptedString []byte) []byte {
	if len(encryptedString) == 0 {
		return []byte{}
	}
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedString[:nonceSize], encryptedString[nonceSize:]
	plainString, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainString
}
