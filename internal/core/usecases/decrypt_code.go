package usecases

import (
	"github.com/MikeMwita/ref-codegen/adapters/output"
)

type DecryptCodeOutput interface {
	PresentDecryptCode(decrypted string)
}

type DecryptCode struct {
	codeService output.CodeService
}

type DecryptCodeInput struct {
	Code string // the code or key to decrypt
	Key  []byte // the key to decrypt with
}

// Execute executes the use case with the given input and output

func (d *DecryptCode) Execute(input DecryptCodeInput, output DecryptCodeOutput) error {
	decrypted, err := d.codeService.DecryptCode(input.Code, input.Key)
	if err != nil {
		return err
	}
	output.PresentDecryptCode(decrypted)
	return nil
}

func NewDecryptCode(codeService output.CodeService) *DecryptCode {
	return &DecryptCode{
		codeService: codeService,
	}
}
