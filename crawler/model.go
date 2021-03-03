package crawler

import (
	"time"
)

// Article qiita の記事を表現する
type Article struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	ID        string    `json:"id"`
	URL       string    `json:"url"`
}
