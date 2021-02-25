package repository

import (
	"context"
	"time"
)

const evalindexcolname = "evaluator.index"

// EvalIndexes EvalIndexesRepository
var EvalIndexes = newEvalIndexes()

type evalIndexes struct {
	*db
}

func newEvalIndexes() *evalIndexes {
	db := newConnection()
	return &evalIndexes{db}
}

// InsertOne Indexの挿入
func (d *evalIndexes) InsertOne(index *Index) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := d.collectionOf(evalindexcolname)
	collection.InsertOne(ctx, index)
}
