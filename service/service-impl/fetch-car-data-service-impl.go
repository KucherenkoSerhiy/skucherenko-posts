package service_impl

import (
	"fmt"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"net/http"
)

type fetchCarDataServiceImpl struct{}

var (
	fetchCarDataUrl = "https://myfakeapi.com/api/cars/1"
)

func NewCarService() service.FetchCarDataService {
	return &fetchCarDataServiceImpl{}
}

func (*fetchCarDataServiceImpl) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the fetchCarDataUrl %s", fetchCarDataUrl)

	resp, _ := client.Get(fetchCarDataUrl)

	carDataChannel <- resp
}
