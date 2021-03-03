package service

import (
	"strings"

	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/model/response"
	"prototype.mathbase.app/tokenizer"
)

// QueryByLatex
func QueryByLatex(query string) ([]*response.Document, error) {

	chunks := strings.Split(query, ",")
	parser := converter.GetParser(converter.Latex)
	tokenizer := tokenizer.MathmlTokenizer{}

	alltoken := []string{}

	for _, chunk := range chunks {
		pse_res, err := parser.Parse(chunk)
		if err != nil {
			return nil, err
		}
		tokens, err := tokenizer.Tokenize(pse_res.Node)
		if err != nil {
			return nil, err
		}
		alltoken = append(alltoken, tokens...)
	}

	return []*response.Document{
		{Title: "gyutaのぶろぐ", URL: "https://mathbase.app"},
		{Title: "ギューたのぶろぐ", URL: "https://mathbase.app"},
	}, nil
}
