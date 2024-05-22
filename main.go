package main

import (
	"crud/handler"
	"net/http"

	_ "crud/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email your@email.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	mux := http.NewServeMux()

	// @Summary Show a hello message
	// @Description Get a hello message
	// @Produce  json
	// @Success 200 {string} string "Hello, World!"
	// @Router /api/v1/hello [get]
	mux.HandleFunc("POST /usuarios", handler.CreateUserHandler)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", httpSwagger.WrapHandler))

	println("go server listening at :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}

}
