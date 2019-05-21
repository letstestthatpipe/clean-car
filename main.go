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

	r.Get("/api/dummy", handlers.DummyApi)
	r.Get("/api/car/{carId}", handlers.GetCarStatus)

	r.Get("/", handleMain)
	r.Get("/login", infra.HandleOauthLogin)
	r.Get("/callback", handlers.HandleToggleCarDoors)

	http.ListenAndServe(":3333", r)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "webapp", http.StatusSeeOther)
}
