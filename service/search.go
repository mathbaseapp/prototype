package service

import (
	"strings"

	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/model/response"
	"prototype.mathbase.app/repository"
	"prototype.mathbase.app/tokenizer"
)

// QueryByLatex Latexで検索
func QueryByLatex(query string) ([]*response.Document, error) {

	chunks := strings.Split(query, ",")
	parser := converter.GetParser(converter.Latex)
	tokenizer := tokenizer.MathmlTokenizer{}

	alltoken := []string{}

	for _, chunk := range chunks {
		pseRes, err := parser.Parse(chunk)
		if err != nil {
			return nil, err
		}
		tokens, err := tokenizer.Tokenize(pseRes.Node)
		if err != nil {
			return nil, err
		}
		alltoken = append(alltoken, tokens...)
	}

	indexes, err := repository.Indexes.SelectSortedIndexes(alltoken)
	if err != nil {
		return nil, err
	}

	documents := []*response.Document{}
	for _, index := range indexes {
		documents = append(documents, &response.Document{Title: index.Title, URL: index.URL})
	}
	return documents, nil
}
