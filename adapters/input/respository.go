// interface_adapters/input/repository.go

package input

import (
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

// CodeRepository is an interface that defines the contract for the code repository
type CodeRepository interface {
	Save(code *entities.Code) error                  // saves a code to the repository
	GetByValue(value string) (*entities.Code, error) // gets a code from the repository by the value
}
