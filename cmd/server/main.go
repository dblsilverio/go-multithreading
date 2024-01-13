package main

import (
	"Multithreading/configs"
	"Multithreading/internal/infra/webserver/handler"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

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

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Printf("Server start failed: %s\n", err)
	}
}
