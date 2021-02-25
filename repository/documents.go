package repository

import (
	"context"
	"time"
)

const colname = "document"

// Documents DocumentRepository
var Documents = newDocuments()

type documents struct {
	*db
}

func newDocuments() *documents {
	db := newConnection()
	return &documents{db}
}

// InsertOne Documentの挿入
func (d documents) InsertOne(doc *Document) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := d.collectionOf(colname)
	collection.InsertOne(ctx, doc)
}
