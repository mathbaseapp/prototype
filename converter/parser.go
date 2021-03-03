package converter

import (
	"encoding/xml"
	"errors"
	"os/exec"

	"prototype.mathbase.app/mathml"
)

// ParseResult パース時の結果を表します
type ParseResult struct {
	Source string
	Node   *mathml.Node
}

// DocumentType ドキュメント形式
type DocumentType string

const (
	// Latex LaTeX
	Latex = DocumentType("latex")
	// MathML mathml
	MathML = DocumentType("mathml")
)

// Parser MathMLに変換するためのパーサーを定義します
type Parser interface {
	Parse(string) (*ParseResult, error)
}

type latexParser struct{}

func (latexParser) Parse(source string) (*ParseResult, error) {
	pandocCmd := "echo '$$" + source + "$$'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	if err != nil {
		return nil, errors.New("pandoc の実行時にエラーが発生しました")
	}
	node := xmlNode{}
	xml.Unmarshal(out, &node)
	mm, err := mathMLFactory(&node)
	if err != nil {
		return nil, err
	}
	return &ParseResult{Source: source, Node: mm}, nil
}

type mathmlParser struct{}

func (mathmlParser) Parse(source string) (*ParseResult, error) {

	bsource := []byte(source)
	node := xmlNode{}
	xml.Unmarshal(bsource, &node)
	mm, err := mathMLFactory(&node)
	if err != nil {
		return nil, err
	}
	return &ParseResult{Source: source, Node: mm}, nil
}

// GetParser 適切なコンテンツパーサーを返却します
func GetParser(docType DocumentType) Parser {

	switch docType {
	case Latex:
		return &latexParser{}
	case MathML:
		return &mathmlParser{}
	default:
		panic("incorrect document type.")
	}
}
