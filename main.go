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

	swaggerDir := filepath.Join(workDir, "swaggerui")
	webappDir := filepath.Join(workDir, "webapp")

	infra.FileServer(r, "/swagger-ui", http.Dir(swaggerDir))
	infra.FileServer(r, "/webapp", http.Dir(webappDir))


	carStatusHandler := handlers.CarStatus{Gateway: &infra.MercedesCarGateway{}}

	r.Get("/api/dummy", carStatusHandler.DummyApi)
	r.Get("/api/car/{carId}", carStatusHandler.GetCarStatus)
	r.Get("/api/car/togglelocks", handlers.HandleToggleCarDoors)

	r.Get("/", handleMain)
	r.Get("/login", infra.HandleOauthLogin)

	http.ListenAndServe(":3333", r)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "webapp", http.StatusSeeOther)
}
