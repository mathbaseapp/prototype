package tokenizer

import (
	"errors"

	"prototype.mathbase.app/mathml"
)

// Tokenizer *mathml.Node を token に分割する
type Tokenizer interface {
	Tokenize(*mathml.Node) ([]string, error)
}

// MathmlTokenizer トークナイザーの実体
type MathmlTokenizer struct {
}

// Tokenize トークナイズ
func (m *MathmlTokenizer) Tokenize(mn *mathml.Node) ([]string, error) {
	slice := mn.Map(func(n *mathml.Node) interface{} {
		return mathml.StringWithNoAttr(n)
	})
	res := make([]string, len(slice))
	for index, item := range slice {
		v, ok := item.(string)
		if !ok {
			return nil, errors.New("トークナイズ時の型アサーションにエラーが発生しました")
		}
		res[index] = v
	}
	return res, nil
}
