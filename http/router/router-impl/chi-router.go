package router_impl

import (
	"fmt"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/http/router"
	"github.com/go-chi/chi"
	"net/http"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() router.Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) DELETE(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (*chiRouter) Serve(port string) error {
	fmt.Printf("Chi HTTP server running on port %v", port)
	return http.ListenAndServe(port, chiDispatcher)
}
