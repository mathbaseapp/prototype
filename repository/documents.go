package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Documents DocumentRepository
var Documents = &documents{newColbase("document")}

type documents struct {
	*colbase
}

// InsertOne Documentの挿入
func (c *documents) InsertOne(doc *Document) (*Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	doc.ID = id.String()
	_, err = c.collection().InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// SelectByURL Documentの取得
func (c *documents) SelectByURL(URL string) (*Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	document := Document{}
	err := c.collection().FindOne(ctx, bson.M{"url": URL}).Decode(&document)
	if err != nil {
		return nil, err
	}
	return &document, nil
}

// SelectByID Documentの取得
func (c *documents) SelectByID(ID string) (*Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	document := Document{}
	err := c.collection().FindOne(ctx, bson.M{"_id": ID}).Decode(&document)
	if err != nil {
		return nil, err
	}
	return &document, nil
}

// StreamAllDocument
func (c *documents) StreamEveryDocument(callback func(Document) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cur, err := c.collection().Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var documents []Document
		err := cur.Decode(&documents)
		if err != nil {
			return err
		}

		for _, doc := range documents {
			err := callback(doc)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}
