package service

import "github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(postId int64) error
}
