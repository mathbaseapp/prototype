package converter

import (
	"encoding/xml"
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
)

// Parser MathMLに変換するためのパーサーを定義します
type Parser interface {
	Parse(string) (ParseResult, error)
}

type latexParser struct{}

// Parse latexparser
func (p latexParser) Parse(source string) (ParseResult, error) {
	pandocCmd := "echo '$$" + source + "$$'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	if err != nil {
		panic("pandoc cannot execute. is not installed?") // TODO エラーハンドリング
	}
	node := xmlNode{}
	xml.Unmarshal(out, &node)
	mNode := mathMLNodeFactory(&node)
	return ParseResult{Source: source, Node: mNode}, nil
}

// GetParser 適切なコンテンツパーサーを返却します
func GetParser(docType DocumentType) Parser {

	return latexParser{}
}
