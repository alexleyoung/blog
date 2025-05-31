package firebase

import (
	"context"
	"log"

	"github.com/alexleyoung/blog/types"
)

func GetPost(ctx context.Context, slug string) (types.Post, error) {
	iter := FirestoreClient.Collection("posts").
		Where("slug", "==", slug).
		Limit(1).
		Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		return types.Post{}, err
	}
	var post types.Post
	doc.DataTo(&post)

	return post, nil
}

func GetAllPosts(ctx context.Context) ([]types.Post, error) {
	postsCollection := FirestoreClient.Collection("posts")

	postDocs, err := postsCollection.Documents(ctx).GetAll()
	if err != nil {
		log.Print("failed to get posts from DB")
		return nil, err
	}

	posts := make([]types.Post, len(postDocs))
	for i, doc := range postDocs {
		var post types.Post
		err := doc.DataTo(&post)
		if err != nil {
			log.Print("failed to convert post to struct")
			continue
		}
		posts[i] = post
	}

	return posts, nil
}

func CreatePost(ctx context.Context, post types.Post) error {
	postsCollection := FirestoreClient.Collection("posts")

	_, _, err := postsCollection.Add(ctx, post)
	if err != nil {
		log.Print("failed to add post to DB")
		return err
	}

	return nil
}
