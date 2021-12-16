package main

import (
	"fmt"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/controller"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/controller/controller-impl"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/http/router"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/http/router/router-impl"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/repository"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/repository/repository-impl"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service/service-impl"
	"net/http"
	"os"
)

var (
	postRepository repository.PostRepository = repository_impl.NewFirestorePostRepository()

	postService service.PostService = service_impl.NewPostService(postRepository)

	postController controller.PostController = controller_impl.NewPostController(postService)

	httpRouter router.Router = router_impl.NewMuxRouter()
)

func main() {
	launch(httpRouter)
}

func launch(router router.Router) {
	router.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})
	router.GET("/posts", postController.GetPosts)
	router.POST("/posts", postController.AddPost)
	router.DELETE("/posts", postController.DeletePost)

	port := ":" + os.Getenv("PORT")
	router.Serve(port)
}
