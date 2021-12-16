package controller_impl

import (
	"encoding/json"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/controller"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/errors"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"log"
	"net/http"
	"strconv"
)

var (
	appService service.PostService
)

type postControllerImpl struct{}

func NewPostController(service service.PostService) controller.PostController {
	appService = service
	return &postControllerImpl{}
}

func (*postControllerImpl) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := appService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
		log.Println(err)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*postControllerImpl) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var post entity.Post
	marshallingError := json.NewDecoder(request.Body).Decode(&post)

	if marshallingError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling posts"})
		log.Println(marshallingError)
		return
	}
	validationError := appService.Validate(&post)
	if validationError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Validation error"})
		log.Println(validationError)
		return
	}

	result, savingError := appService.Create(&post)
	if savingError != nil {
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*postControllerImpl) DeletePost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	postIdParam, hasParameter := request.URL.Query()["id"]
	if !hasParameter {
		log.Println("Missing url param 'id'")
		return
	}

	var postId int64
	postId, parseError := strconv.ParseInt(postIdParam[0], 10, 64)
	if parseError != nil {
		log.Println("Incorrect url param 'id'")
		return
	}

	deletionError := appService.Delete(postId)
	if deletionError != nil {
		return
	}
	response.WriteHeader(http.StatusOK)
}
