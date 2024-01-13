package main

import (
	"Multithreading/configs"
	_ "Multithreading/docs"
	"Multithreading/internal/infra/webserver/handler"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title           Go Expert - CEP Timeout
// @version         1.0
// @description     CEP Search tool.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Diogo Silverio

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	configs, err := configs.LoadConfig("cmd/server/")
	if err != nil {
		panic(err)
	}

	cepHandler := handler.NewCepHandler(configs.Timeout)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", cepHandler.GetCEP)
	})

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/docs/doc.json")))

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Printf("Server start failed: %s\n", err)
	}
}
