package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Documents DocumentRepository
var Documents = &documents{newRepo("document")}

type documents struct {
	*colbase
}

// InsertOne Documentの挿入
func (c *documents) InsertOne(doc *Document) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c.collection().InsertOne(ctx, doc)
}

// SelectByURL Documentの取得
func (c *documents) SelectByURL(URL string) *Document {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	document := Document{}
	c.collection().FindOne(ctx, bson.M{"url": URL}).Decode(&document)
	return &document
}
