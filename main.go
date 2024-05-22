package main

import (
	"crud/handler"
	"net/http"

	_ "crud/docs"

	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost/", "http://localhost:8000/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Máximo de 5 minutos
	})

	mux := http.NewServeMux()

	mux.HandleFunc("POST /usuarios", handler.CreateUserHandler)
	mux.HandleFunc("GET /usuarios", handler.GetUsersHandler)
	mux.HandleFunc("GET /usuarios/{id}", handler.GetUserHandler)

	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))

	handler := c.Handler(mux)
	http.Handle("/", c.Handler(mux))

	println("go server listening at :8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		panic(err)
	}
}
