package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/vbasem/clean-car/handlers"
	"github.com/vbasem/clean-car/infra"
)

func main() {
	log.Printf("starting clean car")
	r := chi.NewRouter()

	workDir, err := os.Getwd()

	if err != nil {
		log.Panicf("got an error: %s", err)
	}

	carStatusHandler := handlers.CarStatus{Gateway: &handlers.MercedesCarGateway{}}

	swaggerDir := filepath.Join(workDir, "swaggerui")

	infra.FileServer(r, "/swagger-ui", http.Dir(swaggerDir))

	r.Get("/api/dummy", carStatusHandler.DummyApi)
	r.Get("/api/car/{carId}", carStatusHandler.GetCarStatus)

	http.ListenAndServe(":3333", r)
}
