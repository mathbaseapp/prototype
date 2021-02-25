package converter

import (
	"encoding/xml"
	"os/exec"
)

// DocumentType ドキュメント形式
type DocumentType string

const (
	// Latex LaTeX
	Latex = DocumentType("latex")
)

// Parser MathMLに変換するためのパーサーを定義します
type Parser interface {
	Parse(string) (MathMLNode, error)
}

type latexParser struct{}

// Parse latexparser
func (p latexParser) Parse(source string) (MathMLNode, error) {
	pandocCmd := "echo '" + source + "'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	if err != nil {
		panic("pandoc cannot execute. is not installed?") // TODO エラーハンドリング
	}
	node := xmlNode{}
	xml.Unmarshal(out, &node)
	return mathMLFactory(&node), nil
}

// GetParser 適切なコンテンツパーサーを返却します
func GetParser(docType DocumentType) Parser {

	return latexParser{}
}
