package handlers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func (CarStatus *CarStatus) DummyApi(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET door status")
	w.Write([]byte("get car status."))
}

type CarStatus struct {
	Gateway CarGateway
}

func (CarStatus *CarStatus) GetCarStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET car status")
	log.Printf("request was %v", r.URL.String())
	carId := chi.URLParam(r, "carId")

	response, err := CarStatus.Gateway.GetCarStatus(carId)

	if err != nil {
		log.Printf("something went wrong with the GET car: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("CAR status response was: %s", response)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
