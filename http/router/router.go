package router

import "net/http"

type Router interface {
	GET(uri string, f func(response http.ResponseWriter, request *http.Request))
	POST(uri string, f func(response http.ResponseWriter, request *http.Request))
	DELETE(uri string, f func(response http.ResponseWriter, request *http.Request))

	Serve(port string) error
}
