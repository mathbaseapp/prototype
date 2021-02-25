package converter

import (
	"encoding/xml"
	"fmt"
	"os/exec"
)

// MarkupStyle マークアップの表現方法を示します
type MarkupStyle string

const (
	// Presentation プレゼンテーションマークアップ
	Presentation = MarkupStyle("Presentation")
	// Content コンテンツマークアップ
	Content = MarkupStyle("Content")
)

// MathMLNode MathMLにおけるNodeを表現します
type MathMLNode interface {
	Name() xml.Name
	Value() string
	Children() []MathMLNode
	Style() MarkupStyle
}

// DefaultMathMLNode 表現されていないNodeが格納されます
type DefaultMathMLNode struct {
	name     xml.Name
	value    string
	children []MathMLNode
}

// Name tagname
func (m DefaultMathMLNode) Name() xml.Name {
	return m.name
}

// Value 値を返す
func (m DefaultMathMLNode) Value() string {
	return m.value
}

// Children 子ノードを返す
func (m DefaultMathMLNode) Children() []MathMLNode {
	return m.children
}

// Style Presentationを返す
func (m DefaultMathMLNode) Style() MarkupStyle {
	return Presentation
}

func mathMLFactory(node *Node) MathMLNode {

	children := []MathMLNode{}
	for _, n := range node.Nodes {
		children = append(children, mathMLFactory(n))
	}

	return DefaultMathMLNode{node.Name, node.Value, children}
}

// Run app
func Run() {
	latexStr := "$$x = 4$$"
	pandocCmd := "echo '" + latexStr + "'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	str := Node{}
	fmt.Println(string(out))
	xml.Unmarshal(out, &str)

	if err != nil {
		fmt.Println("Command Exec Error.")
	}
	fmt.Println(mathMLFactory(&str))
}
