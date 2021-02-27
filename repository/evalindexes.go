package repository

import (
	"context"
	"time"
)

// EvalIndexes EvalIndexesRepository
var EvalIndexes = &evalIndexes{newRepo("evaluator.index")}

type evalIndexes struct {
	*colbase
}

// InsertOne Indexの挿入
func (c *evalIndexes) InsertOne(index *Index) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c.cli().InsertOne(ctx, index)
}
