package repository_impl

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/repository"
	"github.com/KucherenkoSerhiy/skucherenko-posts.git/resources"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
)

type firestorePostRepo struct{}

// NewFirestorePostRepository constructor
func NewFirestorePostRepository() repository.PostRepository {
	return &firestorePostRepo{}
}

func (*firestorePostRepo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, resources.FIREBASE_PROJECT_ID, option.WithCredentialsFile(getCredentialsFileLocation()))
	if err != nil {
		log.Panicf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err2 := client.Collection(resources.FIREBASE_POSTS_COLLECTION_NAME).Add(ctx, map[string]interface{}{
		"ID":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err2 != nil {
		log.Panicf("Failed adding a new post: %v", err2)
		return nil, err2
	}

	return post, nil
}

func (*firestorePostRepo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, resources.FIREBASE_PROJECT_ID, option.WithCredentialsFile(getCredentialsFileLocation()))
	if err != nil {
		log.Panicf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	iter := client.Collection(resources.FIREBASE_POSTS_COLLECTION_NAME).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Panicf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			Id:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// TODO: does not delete
func (*firestorePostRepo) Delete(postId int64) error {
	ctx := context.Background()
	client, clientError := firestore.NewClient(ctx, resources.FIREBASE_PROJECT_ID, option.WithCredentialsFile(getCredentialsFileLocation()))
	if clientError != nil {
		log.Panicf("Failed to create a Firestore Client: %v", clientError)
		return clientError
	}
	defer client.Close()

	_, err := client.Collection(resources.FIREBASE_POSTS_COLLECTION_NAME).Doc(strconv.FormatInt(postId, 10)).Delete(ctx)
	if err != nil {
		log.Panicf("Failed to delete post with id %v. Original error: %v", postId, clientError)
		return err
	}

	return nil
}

func getCredentialsFileLocation() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	credentialsFile := path.Join(basepath, resources.FIREBASE_CREDENTIALS_FILE)
	return credentialsFile
}
