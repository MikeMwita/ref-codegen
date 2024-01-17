// interface_adapters/presenters/text_presenter.go

package presenters

import (
	"fmt"
	"github.com/MikeMwita/ref-codegen/adapters/output"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

// TextPresenter is a type that implements the output interfaces using text format
type TextPresenter struct{}

// NewTextPresenter creates a new TextPresenter
func NewTextPresenter() *TextPresenter {
	return &TextPresenter{}
}

// Present presents a code to the output using text format
func (t *TextPresenter) Present(code *entities.Code) {
	// create a response data structure from the code entity
	response := output.Response{
		Value:  code.Value,
		Format: code.Format,
	}
	// print the response to the output
	fmt.Printf("Code: %s\n", response.Value)
	fmt.Printf("Format: %s\n", response.Format)
}

// Present presents a validation result to the output using text format
func (t *TextPresenter) Present(valid bool) {
	// create a response data structure from the validation result
	response := output.Response{
		Valid: valid,
	}
	// print the response to the output
	fmt.Printf("Valid: %v\n", response.Valid)
}

// Present presents an encrypted or decrypted code to the output using text format
func (t *TextPresenter) Present(code string) {
	// create a response data structure from the code
	response := output.Response{
		Value: code,
	}
	// print the response to the output
	fmt.Printf("Code: %s\n", response.Value)
}
