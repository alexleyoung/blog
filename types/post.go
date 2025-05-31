package types

import "time"

type Post = struct {
	ID    string `firestore:"-"`
	Title string `firestore:"title"`
	Slug  string `firestore:"slug"`
	Body  string `firestore:"body"`

	Tags []string `firestore:"tags"`

	CreatedAt time.Time `firestore:"CreatedAt"`
	UpdatedAt time.Time `firestore:"UpdatedAt"`
}
