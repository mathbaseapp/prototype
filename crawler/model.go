package crawler

import (
	"fmt"
	"time"
)

type article struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	ID        string    `json:"id"`
	URL       string    `json:"url"`
}

type formula struct {
	startLine  int // 何行目に現れたか
	lineLength int // 何行続いたか
	value      []string
}

func (f *formula) getInfo() string {
	var str string
	str += fmt.Sprintf("startLine: %d\t", f.startLine)
	str += fmt.Sprintf("lineLength: %d\n", f.lineLength)
	for _, v := range f.value {
		str += fmt.Sprintf("\t%s\n", v)
	}
	return str
}
