package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Indexes EvalIndexesRepository
var Indexes = &indexes{newColbase("evaluator.index")}

type indexes struct {
	*colbase
}

// InsertOne Indexの挿入
func (c *indexes) InsertOne(index *Index) (*Index, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	index.ID = id.String()
	_, err = c.collection().InsertOne(ctx, index)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func (c *indexes) SelectIndex(keys []string) ([]*Index, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"key": bson.M{"$in": keys},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":      "$document.url",
				"title":    bson.M{"$first": "$document.title"},
				"location": bson.M{"$push": "$location"},
				"count":    bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{"count": -1},
		},
		bson.M{
			"$limit": 30,
		},
	}
	c.collection().Aggregate(ctx, pipeline)
}
