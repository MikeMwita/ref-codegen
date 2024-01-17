// interface_adapters/controllers/cli_controller.go

package controllers

import (
	"fmt"
	"github.com/MikeMwita/ref-codegen/adapters/input"
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
	"os"
)

// CLIController is a type that implements the input interfaces using CLI commands
type CLIController struct {
	generateCode  *usecases.GenerateCode    // the use case for generating codes
	validateCode  *usecases.ValidateCode    // the use case for validating codes
	encryptCode   *usecases.EncryptCode     // the use case for encrypting codes
	decryptCode   *usecases.DecryptCode     // the use case for decrypting codes
	textPresenter *presenters.TextPresenter // the presenter for text format
}

// NewCLIController creates a new CLIController with the given use cases and presenters
func NewCLIController(
	generateCode *usecases.GenerateCode,
	validateCode *usecases.ValidateCode,
	encryptCode *usecases.EncryptCode,
	decryptCode *usecases.DecryptCode,
	textPresenter *presenters.TextPresenter,
) *CLIController {
	return &CLIController{
		generateCode:  generateCode,
		validateCode:  validateCode,
		encryptCode:   encryptCode,
		decryptCode:   decryptCode,
		textPresenter: textPresenter,
	}
}

// GenerateCodeCommand handles the CLI command for generating codes
func (c *CLIController) GenerateCodeCommand(args []string) {
	// parse the arguments from the command
	format := args[0]
	length := args[1]
	prefix := args[2]
	suffix := args[3]
	delimiter := args[4]
	// create a request data structure from the arguments
	request := input.Request{
		Format:    format,
		Length:    length,
		Prefix:    prefix,
		Suffix:    suffix,
		Delimiter: delimiter,
	}
	// convert the request to a GenerateCodeInput
	input := request.ToGenerateCodeInput()
	// choose the presenter for text format
	output := c.textPresenter
	// execute the use case with the input and output
	err := c.generateCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ValidateCodeCommand handles the CLI command for validating codes
func (c *CLIController) ValidateCodeCommand(args []string) {
	// parse the argument from the command
	value := args[0]
	// create a request data structure from the argument
	request := input.Request{
		Value: value,
	}
	// convert the request to a ValidateCodeInput
	input := request.ToValidateCodeInput()
	// choose the presenter for text format
	output := c.textPresenter
	// execute the use case with the input and output
	err := c.validateCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// EncryptCodeCommand handles the CLI command for encrypting codes
func (c *CLIController) EncryptCodeCommand(args []string) {
	// parse the arguments from the command
	value := args[0]
	key := args[1]
	// create a request data structure from the arguments
	request := input.Request{
		Value: value,
		Key:   key,
	}
	// convert the request to a EncryptCodeInput
	input := request.ToEncryptCodeInput()
	// choose the presenter for text format
	output := c.textPresenter
	// execute the use case with the input and output
	err := c.encryptCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// DecryptCodeCommand handles the CLI command for decrypting codes
func (c *CLIController) DecryptCodeCommand(args []string) {
	// parse the arguments from the command
	value := args[0]
	key := args[1]
	// create a request data structure from the arguments
	request := input.Request{
		Value: value,
		Key:   key,
	}
	// convert the request to a DecryptCodeInput
	input := request.ToDecryptCodeInput()
	// choose the presenter for text format
	output := c.textPresenter
	// execute the use case with the input and output
	err := c.decryptCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
