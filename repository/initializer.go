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

var client = genClient()

func genClient() *mongo.Client {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+user+":"+password+"@localhost:27017"))
	if err != nil {
		panic("database initialization failed.")
	}
	return client
}

// colbase datastore
type colbase struct {
	client  *mongo.Client
	colname string
}

func (c colbase) collection() *mongo.Collection {
	return c.client.Database(dbname).Collection(c.colname)
}

func newColbase(colname string) *colbase {
	return &colbase{client: client, colname: colname}
}
