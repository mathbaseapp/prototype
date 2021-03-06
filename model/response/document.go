package response

import "html/template"

// Document ページ情報を格納するレスポンスオブジェクト
type Document struct {
	Title  string        `json:"title"`
	URL    string        `json:"url"`
	Score  float64       `json:"score"`
	MathML template.HTML `json:"mathml"`
}
