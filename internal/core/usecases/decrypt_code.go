// core/usecases/decrypt_code.go

package usecases

// DecryptCode is a use case that decrypts a code or key with the given key
type DecryptCode struct {
	codeService entities.CodeService // the service for encrypting and decrypting codes
}

// NewDecryptCode creates a new DecryptCode use case with the given service
func NewDecryptCode(codeService entities.CodeService) *DecryptCode {
	return &DecryptCode{
		codeService: codeService,
	}
}

// Execute executes the use case with the given input and output
func (d *DecryptCode) Execute(input DecryptCodeInput, output DecryptCodeOutput) error {
	// decrypt the code or key with the given key
	decrypted, err := d.codeService.Decrypt(input.Code, input.Key)
	if err != nil {
		return err
	}
	// present the decrypted code or key to the output
	output.Present(decrypted)
	return nil
}

// DecryptCodeInput is a type that defines the input for the DecryptCode use case
type DecryptCodeInput struct {
	Code string // the code or key to decrypt
	Key  []byte // the key to decrypt with
}

// DecryptCodeOutput is an interface that defines the output for the DecryptCode use case
type DecryptCodeOutput interface {
	Present(decrypted string) // presents the decrypted code or key to the output
}
