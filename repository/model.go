package repository

// Index index
type Index struct {
	Key        string
	FragmentID string
}

// Fragment ページ内の数式
type Fragment struct {
	ID  string
	Loc string
	URL string
}

// Document qiitaのページ
type Document struct {
	URL     string
	Content string
}
