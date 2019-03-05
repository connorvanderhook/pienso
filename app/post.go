package main

// Post data model. I suggest looking at https://upper.io for an easy
// and powerful data persistence adapter.
type Post struct {
	ID string `json:"id"`
	// Author int64  `json:"user_id"` // the author
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"body"`
	Slug  string `json:"slug"`
}
