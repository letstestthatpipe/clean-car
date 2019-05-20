package main

import (
	"github.com/go-chi/chi"
	"github.com/vbasem/clean-car/handlers"
	"github.com/vbasem/clean-car/infra"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.Printf("starting clean car")
	r := chi.NewRouter()

	workDir, err := os.Getwd()

	if err != nil {
		log.Panicf("got an error: %s", err)
	}

	filesDir := filepath.Join(workDir, "swaggerui")

	infra.FileServer(r, "/swagger-ui", http.Dir(filesDir))
	r.Get("/", handlers.redirectToSwagger)
	r.Get("/api/dummy", handlers.DummyApi)
	r.Get("/api/car/{carId}", handlers.GetCarStatus)

	http.ListenAndServe(":3333", r)
}
