// interface_adapters/presenters/json_presenter.go

package presenters

import (
	"encoding/json"
	"github.com/MikeMwita/ref-codegen/adapters/output"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
	"net/http"
)

// JSONPresenter is a type that implements the output interfaces using JSON format
type JSONPresenter struct{}

// NewJSONPresenter creates a new JSONPresenter
func NewJSONPresenter() *JSONPresenter {
	return &JSONPresenter{}
}

// Present presents a code to the output using JSON format
func (j *JSONPresenter) Present(code *entities.Code) {
	// create a response data structure from the code entity
	response := output.Response{
		Value:  code.Value,
		Format: code.Format,
	}
	// encode the response to JSON
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the JSON response to the output
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Present presents a validation result to the output using JSON format
func (j *JSONPresenter) Present(valid bool) {
	// create a response data structure from the validation result
	response := output.Response{
		Valid: valid,
	}
	// encode the response to JSON
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the JSON response to the output
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Present presents an encrypted or decrypted code to the output using JSON format
func (j *JSONPresenter) Present(code string) {
	// create a response data structure from the code
	response := output.Response{
		Value: code,
	}
	// encode the response to JSON
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the JSON response to the output
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
