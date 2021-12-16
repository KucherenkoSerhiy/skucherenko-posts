package router_impl

import (
	"fmt"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/http/router"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() router.Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) DELETE(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) Serve(port string) error {
	fmt.Printf("Mux HTTP server running on port %v", port)
	return http.ListenAndServe(port, muxDispatcher)
}
