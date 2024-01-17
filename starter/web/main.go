package web

import (
	"github.com/MikeMwita/ref-codegen/adapters/controllers"
	"github.com/MikeMwita/ref-codegen/adapters/presenters"
	"github.com/MikeMwita/ref-codegen/internal/core/usecases"
)

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
	// create the presenters
	jsonPresenter := presenters.NewJSONPresenter()
	xmlPresenter := presenters.NewXMLPresenter()
	// create the controller
	controller := controllers.NewHTTPController(
		generateCode,
		validateCode,
		encryptCode,
		decryptCode,
		jsonPresenter,
		xmlPresenter,
	)
	// create the router
	router := NewRouter(controller)
	// start the web server
	router.Run(":8080")
}
