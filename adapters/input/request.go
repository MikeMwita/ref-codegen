package input

import (
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
)

type Request struct {
	Format    string
	Length    int
	Prefix    string
	Suffix    string
	Delimiter string
	Value     string
	Key       []byte
}

//  converts the Request to a GenerateCodeInput

func (r *Request) ToGenerateCodeInput() usecases.GenerateCodeInput {
	return usecases.GenerateCodeInput{
		Format:    r.Format,
		Length:    r.Length,
		Prefix:    r.Prefix,
		Suffix:    r.Suffix,
		Delimiter: r.Delimiter,
	}
}

func (r *Request) ToValidateCodeInput() usecases.ValidateCodeInput {
	return usecases.ValidateCodeInput{
		Value: r.Value,
	}
}

// ToEncryptCodeInput converts the Request to a EncryptCodeInput
func (r *Request) ToEncryptCodeInput() usecases.EncryptCodeInput {
	return usecases.EncryptCodeInput{
		Code: r.Value,
		Key:  r.Key,
	}
}

// converts the Request to a DecryptCodeInput
func (r *Request) ToDecryptCodeInput() usecases.DecryptCodeInput {
	return usecases.DecryptCodeInput{
		Code: r.Value,
		Key:  r.Key,
	}
}
