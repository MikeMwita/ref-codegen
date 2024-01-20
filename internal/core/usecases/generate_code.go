package usecases

import (
	"github.com/MikeMwita/ref-codegen/adapters/input"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

type GenerateCode struct {
	codeRepository input.CodeRepository
}

func (g *GenerateCode) Execute(input GenerateCodeInput, output GenerateCodeOutput) error {
	// create a new Code entity with the given format
	code, err := entities.NewCode("", input.Format)
	if err != nil {
		return err
	}
	// generate a new value for the Code entity with the given options
	err = code.Generate(entities.Options{
		Length:    input.Length,
		Prefix:    input.Prefix,
		Suffix:    input.Suffix,
		Delimiter: input.Delimiter,
	})
	if err != nil {
		return err
	}
	// save the Code entity to the repository
	err = g.codeRepository.Save(code)
	if err != nil {
		return err
	}
	// present the Code entity to the output
	output.Present(code)
	return nil
}

// GenerateCodeInput is a type that defines the input for the GenerateCode use case
type GenerateCodeInput struct {
	Format    string // the format of the code or key, such as Luhn, UUID, KSUID, etc.
	Length    int    // the length of the code or key
	Prefix    string // the prefix of the code or key
	Suffix    string // the suffix of the code or key
	Delimiter string // the delimiter of the code or key
}

// GenerateCodeOutput is an interface that defines the output for the GenerateCode use case
type GenerateCodeOutput interface {
	PresentCode(code *entities.Code) // presents the code or key to the output
}

// NewGenerateCode creates a new GenerateCode use case with the given repository
func NewGenerateCode(codeRepository input.CodeRepository) *GenerateCode {
	return &GenerateCode{
		codeRepository: codeRepository,
	}
}
