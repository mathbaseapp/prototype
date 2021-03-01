package repository

// Index index
type Index struct {
	ID       string        `_id`
	Key      string        `key`
	Location string        `location`
	Document IndexDocument `document`
}

// IndexDocument インデックスに入れるドキュメントを表現します（Documentとの整合性はアプリケーション層で担保して下さい）
type IndexDocument struct {
	ID    string `id`
	URL   string `url`
	Title string `title`
}

// Document qiitaのページ
type Document struct {
	ID      string `_id`
	URL     string `url`
	Title   string `title`
	Content string `content`
}
