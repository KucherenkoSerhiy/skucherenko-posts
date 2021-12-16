package service_impl

import (
	"encoding/json"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"log"
	"net/http"
)

type carDetailsServiceImpl struct{}

var (
	fetchCarDataService service.FetchCarDataService = NewCarService()
	carDataChannel                                  = make(chan *http.Response)

	fetchOwnerDataService service.FetchOwnerDataService = NewOwnerService()
	ownerDataChannel                                    = make(chan *http.Response)
)

func NewCarDetailsService() service.CarDetailsService {
	return &carDetailsServiceImpl{}
}

func (*carDetailsServiceImpl) GetDetails() entity.CarDetails {
	go fetchCarDataService.FetchData()
	go fetchOwnerDataService.FetchData()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		Id:        car.Id,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (entity.Car, error) {
	carData := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(carData.Body).Decode(&car)
	if err != nil {
		log.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	ownerData := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(ownerData.Body).Decode(&owner)
	if err != nil {
		log.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
