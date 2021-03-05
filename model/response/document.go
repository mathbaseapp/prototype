package response

// PageResult ページ情報を格納するレスポンスオブジェクト
type Document struct {
	Title string  `json:"title"`
	URL   string  `json:"url"`
	Point float64 `json:"point"`
}
