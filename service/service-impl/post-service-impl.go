package service_impl

import (
	"errors"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/repository"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/service"
	"math/rand"
)

type postServiceImpl struct{}

var (
	repo repository.PostRepository
)

func NewPostService(postRepository repository.PostRepository) service.PostService {
	repo = postRepository
	return &postServiceImpl{}
}

func (*postServiceImpl) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil
}

func (*postServiceImpl) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	return repo.Save(post)
}

func (*postServiceImpl) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*postServiceImpl) Delete(postId int64) error {
	return repo.Delete(postId)
}
