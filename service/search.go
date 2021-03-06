package service

import (
	"html/template"
	"strings"

	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/model/response"
	"prototype.mathbase.app/repository"
	"prototype.mathbase.app/tokenizer"
)

// QueryByLatex Latexの検索
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
		documents = append(documents, &response.Document{
			Title: index.Title, URL: index.URL, Score: index.Score, MathML: template.HTML(freqEquation(index.Formulas).MathML)})
	}
	return documents, nil
}

func freqEquation(formulas []*repository.FormulaResult) *repository.FormulaResult {

	maxScore := 0.0
	var max *repository.FormulaResult
	for _, formula := range formulas {
		if maxScore <= formula.Score {
			maxScore = formula.Score
			max = formula
		}
	}
	return max
}
