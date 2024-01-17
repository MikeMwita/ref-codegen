// interface_adapters/presenters/xml_presenter.go

package presenters

import (
	"encoding/xml"
	"net/http"
)

// XMLPresenter is a type that implements the output interfaces using XML format
type XMLPresenter struct{}

// NewXMLPresenter creates a new XMLPresenter
func NewXMLPresenter() *XMLPresenter {
	return &XMLPresenter{}
}

// Present presents a code to the output using XML format
func (x *XMLPresenter) Present(code *entities.Code) {
	// create a response data structure from the code entity
	response := output.Response{
		Value:  code.Value,
		Format: code.Format,
	}
	// encode the response to XML
	data, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the XML response to the output
	w.Header().Set("Content-Type", "application/xml")
	w.Write(data)
}

// Present presents a validation result to the output using XML format
func (x *XMLPresenter) Present(valid bool) {
	// create a response data structure from the validation result
	response := output.Response{
		Valid: valid,
	}
	// encode the response to XML
	data, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the XML response to the output
	w.Header().Set("Content-Type", "application/xml")
	w.Write(data)
}

// Present presents an encrypted or decrypted code to the output using XML format
func (x *XMLPresenter) Present(code string) {
	// create a response data structure from the code
	response := output.Response{
		Value: code,
	}
	// encode the response to XML
	data, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the XML response to the output
	w.Header().Set("Content-Type", "application/xml")
	w.Write(data)
}
