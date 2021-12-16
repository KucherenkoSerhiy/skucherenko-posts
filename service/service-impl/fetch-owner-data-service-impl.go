package service_impl

import (
	"fmt"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"net/http"
)

type fetchOwnerDataServiceImpl struct{}

var (
	fetchOwnerDataUrl = "https://myfakeapi.com/api/users/1"
)

func NewOwnerService() service.FetchOwnerDataService {
	return &fetchOwnerDataServiceImpl{}
}

func (*fetchOwnerDataServiceImpl) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the fetchCarDataUrl %s", fetchOwnerDataUrl)

	resp, _ := client.Get(fetchCarDataUrl)

	ownerDataChannel <- resp
}
