package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
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
