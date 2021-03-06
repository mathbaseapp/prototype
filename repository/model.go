package repository

// Index index
type Index struct {
	ID       string        `_id`
	Key      string        `key`
	Weight   float64       `weight`
	Document IndexDocument `document`
	Formula  Formula       `formula`
}

// Formula 数式を表現する
type Formula struct {
	Location int    `location`
	MathML   string `mathml`
}

// IndexDocument インデックスに入れるドキュメントを表現します（Documentとの整合性はアプリケーション層で担保して下さい）
type IndexDocument struct {
	ID    string `id`
	URL   string `url`
	Title string `title`
}

// IndexResult 検索結果を表示
type IndexResult struct {
	URL      string           `_id`
	Title    string           `title`
	Count    int              `count`
	Score    float64          `score`
	Keys     []string         `keys`
	Formulas []*FormulaResult `formula`
}

// FormulaResult 検索結果を表示
type FormulaResult struct {
	Location int     `location`
	MathML   string  `mathml`
	Score    float64 `score`
}

// Document qiitaのページ
type Document struct {
	ID      string `_id`
	URL     string `url`
	Title   string `title`
	Content string `content`
}
