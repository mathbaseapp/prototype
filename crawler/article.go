package crawler

import "time"

type article struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	ID        string    `json:"id"`
	URL       string    `json:"url"`
}
