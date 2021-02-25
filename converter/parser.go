package converter

import (
	"encoding/xml"
	"os/exec"
)

// Parser MathMLに変換するためのパーサーを定義します
type Parser interface {
	Parse(string) (MathMLNode, error)
}

// LatexParser Latex向けパーサー
type LatexParser struct{}

// Parse latexparser
func (p LatexParser) Parse(source string) (MathMLNode, error) {
	pandocCmd := "echo '" + source + "'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	if err != nil {
		panic("pandoc cannot execute. is not installed?") // TODO エラーハンドリング
	}
	node := Node{}
	xml.Unmarshal(out, &node)
	return mathMLFactory(&node), nil
}
