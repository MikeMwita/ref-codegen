package presenters

import (
	"encoding/json"
	"github.com/MikeMwita/ref-codegen/adapters/output"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
	"net/http"
)

// JSONPresenter is a type that implements the output interfaces using JSON format
type JSONPresenter struct{}

func (j *JSONPresenter) Present(code *entities.Code) {
	//TODO implement me
	panic("implement me")
}

// NewJSONPresenter creates a new JSONPresenter
func NewJSONPresenter() *JSONPresenter {
	return &JSONPresenter{}
}

// PresentCode presents a code to the output using JSON format
func (j *JSONPresenter) PresentCode(w http.ResponseWriter, code *entities.Code) {
	response := output.Response{
		Value:  code.Value,
		Format: code.Format,
	}

	j.presentJSON(w, response)
}

// PresentValidation presents a validation result to the output using JSON format
func (j *JSONPresenter) PresentValidation(w http.ResponseWriter, valid bool) {
	response := output.Response{
		Valid: valid,
	}

	j.presentJSON(w, response)
}

// PresentEncrypted presents an encrypted or decrypted code to the output using JSON format
func (j *JSONPresenter) PresentEncrypted(w http.ResponseWriter, code string) {
	response := output.Response{
		Value: code,
	}

	j.presentJSON(w, response)
}

func (j *JSONPresenter) presentJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
