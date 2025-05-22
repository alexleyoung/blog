package types

type Post = struct {
	title string
	slug  string
	body  string

	tags []string

	createdAt time.Time
	updatedAt time.Time
}
