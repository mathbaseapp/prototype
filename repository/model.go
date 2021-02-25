package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

// Index index
type Index struct {
	Key        string `key`
	FragmentID primitive.ObjectID
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
	Content string `content`
}
