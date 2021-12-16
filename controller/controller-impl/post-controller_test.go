package controller_impl

import (
	"bytes"
	"encoding/json"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/controller"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/repository"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/repository/repository-impl"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service/service-impl"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	TITLE string = "Test Title"
	TEXT  string = "Test Text"
)

var (
	postRepository repository.PostRepository = repository_impl.NewFirestorePostRepository()
	postService    service.PostService       = service_impl.NewPostService(postRepository)
	sut            controller.PostController = NewPostController(postService)
)

func TestAddPost(t *testing.T) {
	// setup

	// Create a new HTTP POST request
	var body = []byte(`{"title" : "` + TITLE + `", "text" : "` + TEXT + `"}`)
	request, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(body))

	// Assign HTTP Handler function(controller AddPost function)
	handler := http.HandlerFunc(sut.AddPost)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, request)

	// Add Assertions on the HTTP Status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: expected %v, got %v", http.StatusOK, status)
	}

	// Decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.Id)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// cleanup
}

func TestGetPosts(t *testing.T) {
	request, _ := http.NewRequest("GET", "/posts", nil)

	// Assign HTTP Handler function(controller GetPosts function)
	handler := http.HandlerFunc(sut.GetPosts)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, request)

	// Add Assertions on the HTTP Status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: expected %v, got %v", http.StatusOK, status)
	}

	// Decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.NotNil(t, posts[0].Id)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// cleanup
}
