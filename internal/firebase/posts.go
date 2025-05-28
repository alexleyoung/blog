package firebase

import (
	"context"
	"log"

	"github.com/alexleyoung/blog/types"
)

func GetPosts() []types.Post {
	// TODO:
	return []types.Post{}
}

func CreatePost(post types.Post) error {
	ctx := context.Background()

	postsCollection := FirestoreClient.Collection("posts")

	_, _, err := postsCollection.Add(ctx, post)
	if err != nil {
		log.Print("failed to add post to DB")
		return err
	}

	return nil
}
