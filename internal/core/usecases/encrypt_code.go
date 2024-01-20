package usecases

import (
	"github.com/MikeMwita/ref-codegen/adapters/output"
)

type EncryptCodeOutput interface {
	PresentEncryptCode(encrypted string)
}

type EncryptCode struct {
	codeService output.CodeService
}

// Execute executes the use case with the given input and output
func (e *EncryptCode) Execute(input EncryptCodeInput, output EncryptCodeOutput) error {
	encrypted, err := e.codeService.EncryptCode(input.Code, input.Key)
	if err != nil {
		return err
	}
	output.PresentEncryptCode(encrypted)
	return nil
}

type EncryptCodeInput struct {
	Code string
	Key  []byte
}

func NewEncryptCode(codeService output.CodeService) *EncryptCode {
	return &EncryptCode{
		codeService: codeService,
	}
}
