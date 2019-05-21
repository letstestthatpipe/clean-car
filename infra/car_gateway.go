package infra

import "golang.org/x/oauth2"

type CarGateway interface {
	getCarStatus(carId string) (string, error)
	toggleCarLocks(token oauth2.Token) (string, error)
}

type MercedesCarGateway struct {
}

func (gw *MercedesCarGateway) getCarStatus(carId string) (string, error) {
	return "", nil
}

func toggleCarLocks(token oauth2.Token) (string, error) {
	return "", nil
}
