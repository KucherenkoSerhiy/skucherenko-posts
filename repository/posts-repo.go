package repository

import (
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(postId int64) error
}
