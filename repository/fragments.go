package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const flagcolname = "fragments"

// Fragments FragmentRepository
var Fragments = newFragments()

type fragremts struct {
	*db
}

func newFragments() *fragremts {
	db := newConnection()
	return &fragremts{db}
}

// InsertOne Documentの挿入
func (d *fragremts) InsertOne(frag *Fragment) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := d.collectionOf(flagcolname)
	collection.InsertOne(ctx, frag)
}

// SelectByURL Documentの取得
func (d *fragremts) SelectByID(id primitive.ObjectID) *Fragment {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := d.collectionOf(flagcolname)
	frag := Fragment{}
	collection.FindOne(ctx, bson.M{"_id": id}).Decode(&frag)
	return &frag
}
