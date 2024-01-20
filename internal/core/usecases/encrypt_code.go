// core/usecases/encrypt_code.go

package usecases

import (
	"github.com/MikeMwita/ref-codegen/internal/crypto"
)

// EncryptCode is a use case that encrypts a code or key with the given key
type EncryptCode struct {
	codeService crypto.CodeService // the service for encrypting and decrypting codes
}

// NewEncryptCode creates a new EncryptCode use case with the given service
func NewEncryptCode(codeService crypto.CodeService) *EncryptCode {
	return &EncryptCode{
		codeService: codeService,
	}
}

// Execute executes the use case with the given input and output
func (e *EncryptCode) Execute(input EncryptCodeInput, output EncryptCodeOutput) error {
	// encrypt the code or key with the given key
	encrypted, err := e.codeService.Encrypt(input.Code, input.Key)
	if err != nil {
		return err
	}
	// present the encrypted code or key to the output
	output.Present(encrypted)
	return nil
}

// EncryptCodeInput is a type that defines the input for the EncryptCode use case
type EncryptCodeInput struct {
	Code string // the code or key to encrypt
	Key  []byte // the key to encrypt with
}

// EncryptCodeOutput is an interface that defines the output for the EncryptCode use case
type EncryptCodeOutput interface {
	Present(encrypted string) // presents the encrypted code or key to the output
}
