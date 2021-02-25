package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
func (d *documents) InsertOne(doc *Document) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := d.collectionOf(colname)
	collection.InsertOne(ctx, doc)
}

// SelectByURL Documentの取得
func (d *documents) SelectByURL(URL string) *Document {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := d.collectionOf(colname)
	document := Document{}
	collection.FindOne(ctx, bson.M{"url": URL}).Decode(&document)
	return &document
}
