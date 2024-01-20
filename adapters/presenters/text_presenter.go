package presenters

import (
	"fmt"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

type TextPresenter struct{}

func (t *TextPresenter) PresentCode(code *entities.Code) {
	fmt.Printf("Code: %s\n", code.Value)
}

func (t *TextPresenter) PresentValidationResult(valid bool) {
	fmt.Printf("Validation Result: %t\n", valid)
}

func (t *TextPresenter) PresentEncryptedOrDecryptedCode(code string) {
	fmt.Printf("Encrypted or Decrypted Code: %s\n", code)
}

func (t *TextPresenter) PresentEncryptCode(encrypted string) {
	fmt.Printf("Encrypted Code: %s\n", encrypted)
}

func (t *TextPresenter) PresentDecryptCode(decrypted string) {
	fmt.Printf("Decrypted Code: %s\n", decrypted)
}

func NewTextPresenter() *TextPresenter {
	return &TextPresenter{}
}
