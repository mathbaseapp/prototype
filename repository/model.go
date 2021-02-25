package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

// Index index
type Index struct {
	Key        string             `key`
	FragmentID primitive.ObjectID `fragment_id`
}

// Fragment ページ内の数式
type Fragment struct {
	ID  primitive.ObjectID `_id`
	Loc string             `loc`
	URL string             `url`
}

// Document qiitaのページ
type Document struct {
	URL     string `url`
	Title   string `title`
	Content string `content`
}
