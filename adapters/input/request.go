// interface_adapters/input/request.go

package input

import (
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
)

// Request is a type that defines the input data structures for the interface adapters layer
type Request struct {
	Format    string // the format of the code or key, such as Luhn, UUID, KSUID, etc.
	Length    int    // the length of the code or key
	Prefix    string // the prefix of the code or key
	Suffix    string // the suffix of the code or key
	Delimiter string // the delimiter of the code or key
	Value     string // the value of the code or key
	Key       []byte // the key to encrypt or decrypt with
}

// ToGenerateCodeInput converts the Request to a GenerateCodeInput
func (r *Request) ToGenerateCodeInput() usecases.GenerateCodeInput {
	return usecases.GenerateCodeInput{
		Format:    r.Format,
		Length:    r.Length,
		Prefix:    r.Prefix,
		Suffix:    r.Suffix,
		Delimiter: r.Delimiter,
	}
}

// ToValidateCodeInput converts the Request to a ValidateCodeInput
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

// ToDecryptCodeInput converts the Request to a DecryptCodeInput
func (r *Request) ToDecryptCodeInput() usecases.DecryptCodeInput {
	return usecases.DecryptCodeInput{
		Code: r.Value,
		Key:  r.Key,
	}
}
