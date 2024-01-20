package controllers

import (
	"fmt"
	"github.com/MikeMwita/ref-codegen/adapters/input"
	"github.com/MikeMwita/ref-codegen/adapters/presenters"
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
	"os"
	"strconv"
)

type CLIController struct {
	generateCode  *usecases.GenerateCode
	validateCode  *usecases.ValidateCode
	encryptCode   *usecases.EncryptCode
	decryptCode   *usecases.DecryptCode
	textPresenter *presenters.TextPresenter
}

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

// GenerateCodeCommand is a CLI command for generating a code
func (c *CLIController) GenerateCodeCommand(args []string) {
	if len(args) != 5 {
		fmt.Println("Usage: generate <format> <length> <prefix> <suffix> <delimiter>")
		os.Exit(1)
	}

	format := args[0]
	lengthStr := args[1]
	prefix := args[2]
	suffix := args[3]
	delimiter := args[4]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		fmt.Println("Invalid length argument:", lengthStr)
		os.Exit(1)
	}

	request := input.Request{
		Format:    format,
		Length:    length,
		Prefix:    prefix,
		Suffix:    suffix,
		Delimiter: delimiter,
	}

	input := request.ToGenerateCodeInput()

	output := c.textPresenter

	err = c.generateCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ValidateCodeCommand handles the CLI command for validating codes
func (c *CLIController) ValidateCodeCommand(args []string) {
	value := args[0]
	request := input.Request{
		Value: value,
	}
	input := request.ToValidateCodeInput()
	output := c.textPresenter
	err := c.validateCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// EncryptCodeCommand handles the CLI command for encrypting codes
func (c *CLIController) EncryptCodeCommand(args []string) {
	value := args[0]
	key := args[1]
	request := input.Request{
		Value: value,
		Key:   []byte(key),
	}
	input := request.ToEncryptCodeInput()
	output := c.textPresenter
	err := c.encryptCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// DecryptCodeCommand handles the CLI command for decrypting codes
func (c *CLIController) DecryptCodeCommand(args []string) {
	value := args[0]
	key := args[1]
	request := input.Request{
		Value: value,
		Key:   []byte(key),
	}
	input := request.ToDecryptCodeInput()
	output := c.textPresenter
	err := c.decryptCode.Execute(input, output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
