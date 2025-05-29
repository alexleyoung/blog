package firebase

import (
	"context"
	"log"

	"github.com/alexleyoung/blog/types"
)

func GetAllPosts() ([]types.Post, error) {
	postsCollection := FirestoreClient.Collection("posts")

	postSnapshots, err := postsCollection.Documents(context.Background()).GetAll()
	if err != nil {
		log.Print("failed to get posts from DB")
		return nil, err
	}

	posts := make([]types.Post, len(postSnapshots))
	for i, snapshot := range postSnapshots {
		var post types.Post
		err := snapshot.DataTo(&post)
		if err != nil {
			log.Print("failed to convert post to struct")
			continue
		}
		posts[i] = post
	}

	return posts, nil
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
