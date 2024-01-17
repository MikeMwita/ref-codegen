// frameworks/crypto/crypto.go

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// CodeService is a type that implements the code service interface using AES encryption and decryption
type CodeService struct{}

// NewCodeService creates a new CodeService
func NewCodeService() *CodeService {
	return &CodeService{}
}

// Encrypt encrypts a code with a key using AES
func (c *CodeService) Encrypt(code string, key []byte) (string, error) {
	// convert the code to bytes
	plaintext := []byte(code)
	// create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// create a byte slice for the ciphertext
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// create a byte slice for the initialization vector
	iv := ciphertext[:aes.BlockSize]
	// fill the initialization vector with random data
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	// create a stream cipher with the block and the initialization vector
	stream := cipher.NewCFBEncrypter(block, iv)
	// encrypt the plaintext and store it in the ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// encode the ciphertext to base64
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	// return the encoded string
	return encoded, nil
}

// Decrypt decrypts a code with a key using AES
func (c *CodeService) Decrypt(code string, key []byte) (string, error) {
	// decode the code from base64
	ciphertext, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}
	// create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// check the length of the ciphertext
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	// get the initialization vector from the ciphertext
	iv := ciphertext[:aes.BlockSize]
	// get the actual ciphertext from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]
	// create a byte slice for the plaintext
	plaintext := make([]byte, len(ciphertext))
	// create a stream cipher with the block and the initialization vector
	stream := cipher.NewCFBDecrypter(block, iv)
	// decrypt the ciphertext and store it in the plaintext
	stream.XORKeyStream(plaintext, ciphertext)
	// convert the plaintext to string
	decoded := string(plaintext)
	// return the decoded string
	return decoded, nil
}
