package usecases

import (
	"github.com/MikeMwita/ref-codegen/adapters/input"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

type GenerateCodeOutput interface {
	PresentCode(code *entities.Code)
}

type GenerateCode struct {
	codeRepository input.CodeRepository
}

type GenerateCodeInput struct {
	Format    string // the format of the code or key, such as Luhn, UUID, KSUID, etc.
	Length    int
	Prefix    string
	Suffix    string
	Delimiter string
}

func (g *GenerateCode) Execute(input GenerateCodeInput, output GenerateCodeOutput) error {
	code, err := entities.NewCode("", input.Format)
	if err != nil {
		return err
	}
	err = code.Generate(entities.Options{
		Length:    input.Length,
		Prefix:    input.Prefix,
		Suffix:    input.Suffix,
		Delimiter: input.Delimiter,
	})
	if err != nil {
		return err
	}
	err = g.codeRepository.Save(code)
	if err != nil {
		return err
	}
	output.PresentCode(code)
	return nil
}

func NewGenerateCode(codeRepository input.CodeRepository) *GenerateCode {
	return &GenerateCode{
		codeRepository: codeRepository,
	}
}
