package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
)

func DummyApi(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET door status")
	w.Write([]byte("get car status."))
}


func GetCarStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET car")
	log.Printf("request was %v", r.URL.String())
	carId := chi.URLParam(r, "carId")

	request, _ := http.NewRequest("GET",
		fmt.Sprintf("https://api.mercedes-benz.com/vehicledata_tryout/v1/vehicles/%s/containers/vehiclestatus", carId),
		nil)

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
}

func RedirectToSwagger(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "swagger-ui", http.StatusSeeOther)
}
