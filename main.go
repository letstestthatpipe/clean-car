package main

import (
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() { 
	log.Printf("starting clean car")
	r := chi.NewRouter()


	workDir, err := os.Getwd()

	if err != nil { 
		log.Panicf("got an error: %s", err)
	}

	filesDir := filepath.Join(workDir, "swaggerui")

	FileServer(r, "/swagger-ui", http.Dir(filesDir))
	CarClient(r)

	http.ListenAndServe(":3333", r)


}

func CarClient(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET door status")
		w.Write([]byte("get door status."))
	})

	r.Get("/car", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET car")
		request, _ := http.NewRequest("GET", "https://api.mercedes-benz.com/vehicledata_tryout/v1/vehicles/WDB111111ZZZ22222/containers/vehiclestatus", nil)
		request.Header.Set("Authorization", "Bearer 4c4c444c-v123-4123-s123-4c4c444c4c44")

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			log.Printf("something went wrong with the GET car: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Printf("something went wrong with the GET car: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		

		log.Printf("CAR status response was: %s", body)
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}