package main

import (
	"fmt"
	"github.com/MikeMwita/ref-codegen/adapters/controllers"
	"github.com/MikeMwita/ref-codegen/adapters/presenters"
	"github.com/MikeMwita/ref-codegen/database"
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
	"github.com/MikeMwita/ref-codegen/internal/crypto"
	"os"
)

type Commands struct {
	Controller *controllers.CLIController
}

func NewCommands() *Commands {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	codeRepository := database.NewCodeRepository(db)

	//  code service
	codeService := crypto.NewCodeService()

	//  use cases
	generateCode := usecases.NewGenerateCode(codeRepository)
	validateCode := usecases.NewValidateCode(codeRepository)
	encryptCode := usecases.NewEncryptCode(codeService)
	decryptCode := usecases.NewDecryptCode(codeService)

	textPresenter := presenters.NewTextPresenter()

	controller := controllers.NewCLIController(
		generateCode,
		validateCode,
		encryptCode,
		decryptCode,
		textPresenter,
	)

	return &Commands{
		Controller: controller,
	}
}

func (c *Commands) GenerateCommand(args []string) {
	c.Controller.GenerateCodeCommand(args)
}

func (c *Commands) ValidateCommand(args []string) {
	c.Controller.ValidateCodeCommand(args)
}

func (c *Commands) EncryptCommand(args []string) {
	c.Controller.EncryptCodeCommand(args)
}

func (c *Commands) DecryptCommand(args []string) {
	// Call the corresponding handler in the controller
	c.Controller.DecryptCodeCommand(args)
}

func UnknownCommand() {
	fmt.Println("Unknown command. Use 'help' for available commands.")
	os.Exit(1)
}
