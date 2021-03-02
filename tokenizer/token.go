package tokenizer

import (
	"errors"

	"prototype.mathbase.app/mathml"
	"prototype.mathbase.app/middlelng"
)

// Tokenizer middleLanguage を token に分割する
type Tokenizer interface {
	Tokenize(middlelng.MiddleLanguage) ([]string, error)
}

// MathmlTokenizer mathml専用のトークナイザー
type MathmlTokenizer struct {
}

// Tokenize トークナイズ
func (m *MathmlTokenizer) Tokenize(ml middlelng.MiddleLanguage) ([]string, error) {
	slice := ml.Map(func(n middlelng.MiddleLanguage) interface{} {
		return mathml.Printer(n.(*mathml.Node))
	})
	res := make([]string, len(slice))
	for index, item := range slice {
		v, ok := item.(string)
		if !ok {
			return make([]string, 0), errors.New("トークナイズ時の型アサーションにエラーが発生しました")
		}
		res[index] = v
	}
	return res, nil
}
