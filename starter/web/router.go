// frameworks/web/router.go

package main

import (
	"github.com/MikeMwita/ref-codegen/adapters/controllers"
)

// NewRouter creates a new router with the given controller
func NewRouter(controller *controllers.HTTPController) *gin.Engine {
	// create a new gin engine
	router := gin.Default()
	// define the routes
	router.GET("/generate", controller.GenerateCodeHandler)
	router.GET("/validate", controller.ValidateCodeHandler)
	router.GET("/encrypt", controller.EncryptCodeHandler)
	router.GET("/decrypt", controller.DecryptCodeHandler)
	// return the router
	return router
}
