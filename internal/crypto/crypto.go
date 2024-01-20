package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type CodeService struct{}

func (c *CodeService) EncryptCode(code string, key []byte) (string, error) {
	plaintext := []byte(code)
	// create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

func (c *CodeService) DecryptCode(code string, key []byte) (string, error) {
	// decode the code from base64
	ciphertext, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]

	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))

	// create a stream cipher with the block and the initialization vector
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	decoded := string(plaintext)
	return decoded, nil
}

func NewCodeService() *CodeService {
	return &CodeService{}
}
