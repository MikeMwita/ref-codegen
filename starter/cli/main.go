// frameworks/cli/main.go

package main

import (
	"flag"


func main() {
	// create the database connection
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	// create the code repository
	codeRepository := database.NewCodeRepository(db)
	// create the code service
	codeService := crypto.NewCodeService()
	// create the use cases
	generateCode := usecases.NewGenerateCode(codeRepository)
	validateCode := usecases.NewValidateCode(codeRepository)
	encryptCode := usecases.NewEncryptCode(codeService)
	decryptCode := usecases.NewDecryptCode(codeService)
	// create the presenter
	textPresenter := presenters.NewTextPresenter()
	// create the controller
	controller := controllers.NewCLIController(
		generateCode,
		validateCode,
		encryptCode,
		decryptCode,
		textPresenter,
	)
	// parse the command line arguments
	flag.Parse()
	args := flag.Args()
	// get the command name and arguments
	command := args[0]
	args = args[1:]
