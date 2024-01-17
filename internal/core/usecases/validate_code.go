// core/usecases/validate_code.go

package usecases

// ValidateCode is a use case that validates a code or key
type ValidateCode struct {
	codeRepository entities.CodeRepository // the repository for storing and retrieving codes
}

// NewValidateCode creates a new ValidateCode use case with the given repository
func NewValidateCode(codeRepository entities.CodeRepository) *ValidateCode {
	return &ValidateCode{
		codeRepository: codeRepository,
	}
}

// Execute executes the use case with the given input and output
func (v *ValidateCode) Execute(input ValidateCodeInput, output ValidateCodeOutput) error {
	// get the Code entity from the repository by the value
	code, err := v.codeRepository.GetByValue(input.Value)
	if err != nil {
		return err
	}
	// validate the Code entity
	valid, err := code.Validate()
	if err != nil {
		return err
	}
	// present the validation result to the output
	output.Present(valid)
	return nil
}

// ValidateCodeInput is a type that defines the input for the ValidateCode use case
type ValidateCodeInput struct {
	Value string // the value of the code or key
}

// ValidateCodeOutput is an interface that defines the output for the ValidateCode use case
type ValidateCodeOutput interface {
	Present(valid bool) // presents the validation result to the output
}
