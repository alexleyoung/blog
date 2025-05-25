package types

import "time"

type Post = struct {
	Title string
	Slug  string
	Body  string

	Tags []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

var MockPosts = []Post{
	{
		Title:     "Hello World",
		Slug:      "hello-world",
		Body:      "This is my first post!",
		Tags:      []string{"hello", "world"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Title:     "My Second Post",
		Slug:      "my-second-post",
		Body:      "This is my second post!",
		Tags:      []string{"second", "post"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
