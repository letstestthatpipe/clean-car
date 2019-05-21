package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vbasem/clean-car/car"
	"github.com/vbasem/clean-car/infra"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleToggleCarDoors(w http.ResponseWriter, r *http.Request) {
	token, err := infra.GetToken(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Printf("token: %v\n", token)

	getMyCars(w, r, token)
}

func getMyCars(w http.ResponseWriter, r *http.Request, token *oauth2.Token) {
	request, _ := http.NewRequest("GET",
		"https://api.mercedes-benz.com/experimental/connectedvehicle/v1/vehicles",
		nil)

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("something went wrong get my car: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("something went wrong get my car: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var car []car.Car
	err = json.Unmarshal(body, &car)

	if err != nil {
		log.Printf("something went wrong get my car: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("got cars %v", car)
	getMyDoors(w, r, token, car[0].Id)
}

func getMyDoors(w http.ResponseWriter, r *http.Request, token *oauth2.Token, carId string) {

	request, _ := http.NewRequest("GET",
		fmt.Sprintf("https://api.mercedes-benz.com/experimental/connectedvehicle/v1/vehicles/%s/doors", carId),
		nil)

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("something went wrong with getMyDoors: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("something went wrong with getMyDoors: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var doors car.CarDoors
	err = json.Unmarshal(body, &doors)

	if err != nil {
		log.Printf("something went wrong with getMyDoors: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("got door status %v", doors)
	toggleLock(w, r, token, carId, &doors)
}

func toggleLock(w http.ResponseWriter, r *http.Request, token *oauth2.Token, carId string, doors *car.CarDoors) {
	command := doors.GetToggleLocksCommand()
	log.Printf("command will be: %s", command)
	request, _ := http.NewRequest("POST",
		fmt.Sprintf("https://api.mercedes-benz.com/experimental/connectedvehicle/v1/vehicles/%s/doors", carId),
		bytes.NewBuffer([]byte(command)))

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("something went wrong with toggleLock: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("something went wrong with toggleLock: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("got response for changing door status %v", string(body))
	w.Write(body)
}