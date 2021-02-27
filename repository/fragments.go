package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fragments FragmentRepository
var Fragments = &fragremts{newRepo("fragments")}

type fragremts struct {
	*colbase
}

// InsertOne Documentの挿入
func (c *fragremts) InsertOne(frag *Fragment) (*Fragment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := c.cli().InsertOne(ctx, frag)
	if err != nil {
		return nil, err
	}
	frag.ID = res.InsertedID.(primitive.ObjectID)
	return frag, nil
}

// SelectByURL Documentの取得
func (c *fragremts) SelectByID(id primitive.ObjectID) *Fragment {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	frag := Fragment{}
	c.cli().FindOne(ctx, bson.M{"_id": id}).Decode(&frag)
	return &frag
}
