package repository

// Index index
type Index struct {
	Key        string
	FragmentID string
}

// Fragment ページ内の数式
type Fragment struct {
	ID         string
	DocumentID string
	Loc        string
	URL        string // 正規化崩そう
}

// Document qiitaのページ
type Document struct {
	ID      string
	URL     string
	Content string
}
