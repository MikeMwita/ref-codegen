package input

import (
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

type CodeRepository interface {
	Save(code *entities.Code) error
	GetByValue(value string) (*entities.Code, error)
}
