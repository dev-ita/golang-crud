package main

import (
	"crud/handler"
	"net/http"

	_ "crud/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Documentação para o Projeto Integrado II
// @version 1.0
// @description Isso é uma simples aplicação CRUD desenvolvida em GO
// @termsOfService http://swagger.io/terms/

// @contact.name Ítalo Oliveira
// @contact.url https://github.com/dev-ita
// @contact.email italo.ods@hotmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /usuarios", handler.CreateUserHandler)

	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))

	println("go server listening at :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}

}
