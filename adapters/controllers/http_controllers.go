// interface_adapters/controllers/http_controller.go

package controllers

import (
	"github.com/MikeMwita/ref-codegen/adapters/input"
	"github.com/MikeMwita/ref-codegen/adapters/presenters"
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
	"net/http"

)

// HTTPController is a type that implements the input interfaces using HTTP handlers

type HTTPController struct {
	generateCode  *usecases.GenerateCode  // the use case for generating codes
	validateCode  *usecases.ValidateCode  // the use case for validating codes
	encryptCode   *usecases.EncryptCode   // the use case for encrypting codes
	decryptCode   *usecases.DecryptCode   // the use case for decrypting codes
	jsonPresenter *presenters.JSONPresenter // the presenter for JSON format
	xmlPresenter  *presenters.XMLPresenter  // the presenter for XML format
}

// NewHTTPController creates a new HTTPController with the given use cases and presenters
func NewHTTPController(
	generateCode *usecases.GenerateCode,
	validateCode *usecases.ValidateCode,
	encryptCode *usecases.EncryptCode,
	decryptCode *usecases.DecryptCode,
	jsonPresenter *presenters.JSONPresenter,
	xmlPresenter *presenters.XMLPresenter,
) *HTTPController {
	return &HTTPController{
		generateCode:  generateCode,
		validateCode:  validateCode,
		encryptCode:   encryptCode,
		decryptCode:   decryptCode,
		jsonPresenter: jsonPresenter,
		xmlPresenter:  xmlPresenter,
	}
}

// GenerateCodeHandler handles the HTTP request for generating codes
func (h *HTTPController) GenerateCodeHandler(w http.ResponseWriter, r *http.Request) {
	// parse the query parameters from the request
	format := r.URL.Query().Get("format")
	length := r.URL.Query().Get("length")
	prefix := r.URL.Query().Get("prefix")
	suffix := r.URL.Query().Get("suffix")
	delimiter := r.URL.Query().Get("delimiter")
	// create a request data structure from the query parameters
	request := input.Request{
		Format:    format,
		Length:    length,
		Prefix:    prefix,
		Suffix:    suffix,
		Delimiter: delimiter,
	}
	// convert the request to a GenerateCodeInput
	input := request.ToGenerateCodeInput()
	// get the accept header from the request
	accept := r.Header.Get("Accept")
	// choose the presenter based on the accept header
	var output usecases.GenerateCodeOutput
	switch accept {
	case "application/json":
		output = h.jsonPresenter
	case "application/xml":
		output = h.xmlPresenter
	default:
		http.Error(w, "Unsupported format", http.StatusBadRequest)
		return
	}
	// execute the use case with the input and output
	err := h.generateCode.Execute(input, output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ValidateCodeHandler handles the HTTP request for validating codes
func (h *HTTPController) ValidateCodeHandler(w http.ResponseWriter, r *http.Request) {
	// parse the query parameter from the request
	value := r.URL.Query().Get("value")
	// create a request data structure from the query parameter
	request := input.Request{
		Value: value,
	}
	// convert the request to a ValidateCodeInput
	input := request.ToValidateCodeInput()
	// get the accept header from the request
	accept := r.Header.Get("Accept")
	// choose the presenter based on the accept header
	var output usecases.ValidateCodeOutput
	switch accept {
	case "application/json":
		output = h.jsonPresenter
	case "application/xml":
		output = h.xmlPresenter
	default:
		http.Error(w, "Unsupported format", http.StatusBadRequest)
		return
	}
	// execute the use case with the input and output
	err := h.validateCode.Execute(input, output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// EncryptCodeHandler handles the HTTP request for encrypting codes
func (h *HTTPController) EncryptCodeHandler(w http.ResponseWriter, r *http.Request) {
	// parse the query parameters from the request
	value := r.URL.Query().Get("value")
	key := r.URL.Query().Get("key")
	// create a request data structure from the query parameters
	request := input.Request{
		Value: value,
		Key:   key,
	}
	// convert the request to a EncryptCodeInput
	input := request.ToEncryptCodeInput()
	// get the accept header from the request
	accept := r.Header.Get("Accept")
	// choose the presenter based on the accept header
	var output usecases