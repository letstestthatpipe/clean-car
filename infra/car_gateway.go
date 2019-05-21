package infra

import (
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

type CarGateway interface {
    GetCarStatus(carId string) (string, error)
	ToggleCarLocks(token oauth2.Token) (string, error)
}

type MercedesCarGateway struct {
}

func (gw *MercedesCarGateway) GetCarStatus(carId string) (string, error) {
	request, _ := http.NewRequest("GET",
		fmt.Sprintf("https://api.mercedes-benz.com/vehicledata_tryout/v1/vehicles/%s/containers/vehiclestatus", carId),
		nil)

	request.Header.Set("Authorization", "Bearer 4c4c444c-v123-4123-s123-4c4c444c4c44")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (gw *MercedesCarGateway) ToggleCarLocks(token oauth2.Token) (string, error) {
	return "", nil
}
