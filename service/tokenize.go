package service

import (
	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/repository"
	"prototype.mathbase.app/tokenizer"
)

// Tokenize db内のドキュメントに対してトークナイズする
func Tokenize() {

	parser := converter.GetParser(converter.Latex)
	tokenizer := &tokenizer.MathmlTokenizer{}
	processor := &QiitaArticleProcessor{Parser: parser, Tokenizer: tokenizer}

	repository.Documents.StreamEveryDocument(processor.Process)
}
