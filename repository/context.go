package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	user     = "root"
	password = "passwd"
	dbname   = "mathbase"
)

// db datastore
type db struct {
	client *mongo.Client
}

func (d db) collectionOf(colname string) *mongo.Collection {
	return d.client.Database(dbname).Collection(colname)
}

// newConnection create datastore
func newConnection() *db {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+user+":"+password+"@localhost:27017"))
	if err != nil {
		panic("database initialization failed.")
	}
	return &db{client}
}
