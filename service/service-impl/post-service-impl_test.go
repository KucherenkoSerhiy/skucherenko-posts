package service_impl

import (
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestValidateEmptyPost(t *testing.T) {
	sut := NewPostService(nil)

	err := sut.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "T"}
	sut := NewPostService(nil)

	err := sut.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestValidateValidPost(t *testing.T) {
	post := entity.Post{Id: 1, Title: "T", Text: "Text"}
	sut := NewPostService(nil)

	err := sut.Validate(&post)

	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	post := entity.Post{Id: 1, Title: "T", Text: "Text"}
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)
	sut := NewPostService(mockRepo)

	result, err := sut.FindAll()

	mockRepo.AssertExpectations(t)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, post, result[0])
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	post := entity.Post{Id: 1, Title: "T", Text: "Text"}
	mockRepo := new(MockRepository)
	mockRepo.On("Save").Return(&post, nil)
	sut := NewPostService(mockRepo)

	result, err := sut.Create(&post)

	mockRepo.AssertExpectations(t)
	assert.Same(t, &post, result)
	assert.Nil(t, err)
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) Delete(postId int64) error {
	panic("implement me")
}
