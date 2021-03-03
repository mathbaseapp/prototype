package converter

import (
	"encoding/xml"
	"os/exec"
	"regexp"

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

var noutf8 = regexp.MustCompile(`&.*;`) // non-greedy

type latexParser struct{}

func (l *latexParser) Parse(source string) (*ParseResult, error) {
	l.panicIfNoDependency()
	pandocCmd := "echo '$$" + source + "$$'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	if err != nil {
		return nil, err
	}
	uxml := noutf8.ReplaceAllString(string(out), "") // utf-8に含まれない文字/実体参照を削除
	node := &xmlNode{}
	err = xml.Unmarshal([]byte(uxml), node)
	if err != nil {
		return nil, err
	}
	mm, err := mathMLFactory(node)
	if err != nil {
		return nil, err
	}
	return &ParseResult{Source: source, Node: mm}, nil
}

func (latexParser) panicIfNoDependency() {
	if _, err := exec.Command("pandoc", "-v").Output(); err != nil {
		panic("pandoc cannot execute. is not installed?")
	}
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
