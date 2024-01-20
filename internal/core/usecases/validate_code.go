package usecases

import "github.com/MikeMwita/ref-codegen/adapters/input"

type ValidateCodeOutput interface {
	PresentValidationResult(valid bool)
}

type ValidateCode struct {
	codeRepository input.CodeRepository
}

type ValidateCodeInput struct {
	Value string // the value of the code or key
}

func (v *ValidateCode) Execute(input ValidateCodeInput, output ValidateCodeOutput) error {
	code, err := v.codeRepository.GetByValue(input.Value)
	if err != nil {
		return err
	}
	valid, err := code.Validate()
	if err != nil {
		return err
	}

	output.PresentValidationResult(valid)
	return nil
}

func NewValidateCode(codeRepository input.CodeRepository) *ValidateCode {
	return &ValidateCode{
		codeRepository: codeRepository,
	}
}
