package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/vbasem/clean-car/infra"
)

func main() {
	log.Printf("starting clean car")
	r := chi.NewRouter()

	workDir, err := os.Getwd()

	if err != nil {
		log.Panicf("got an error: %s", err)
	}

	swaggerDir := filepath.Join(workDir, "swaggerui")

	infra.FileServer(r, "/swagger-ui", http.Dir(swaggerDir))

	http.ListenAndServe(":3333", r)
}
