package controller

import "net/http"

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
	DeletePost(response http.ResponseWriter, request *http.Request)
}
