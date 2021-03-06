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

func (c *indexes) InsertMany(indexes []*Index) ([]*Index, error) {
	if len(indexes) == 0 {
		return indexes, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for _, index := range indexes {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		index.ID = id.String()
	}
	var anys []interface{}
	for _, any := range indexes {
		anys = append(anys, any)
	}
	_, err := c.collection().InsertMany(ctx, anys)
	if err != nil {
		return nil, err
	}
	return indexes, nil
}

func (c *indexes) SelectSortedIndexes(keys []string) ([]*IndexResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"key": bson.M{"$in": keys},
			},
		},
		{
			"$group": bson.M{
				"_id":      "$document.url",
				"title":    bson.M{"$first": "$document.title"},
				"location": bson.M{"$push": "$location"},
				"count":    bson.M{"$sum": 1},
				"keys":     bson.M{"$push": "$key"},
				"eval":     bson.M{"$sum": "$weight"},
			},
		},
		{
			"$sort": bson.M{"eval": -1},
		},
		{
			"$limit": 20,
		},
	}
	csr, err := c.collection().Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	results := []*IndexResult{}
	err = csr.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (c *indexes) SelectByID(ID string) (*Index, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	index := Index{}
	err := c.collection().FindOne(ctx, bson.M{"document.id": ID}).Decode(&index)
	if err != nil {
		return nil, err
	}

	return &index, err
}
