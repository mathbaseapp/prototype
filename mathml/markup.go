package mathml

import "encoding/xml"

// DefaultNode nodeを返します
func DefaultNode(name xml.Name, value string, children []Node, style MarkupStyle) Node {
	return defaultNode{name: name, value: value, children: children, style: style}
}

// defaultNode 表現されていないNodeが格納されます
type defaultNode struct {
	name     xml.Name
	value    string
	children []Node
	style    MarkupStyle
}

// GetName tagname
func (m defaultNode) GetName() xml.Name {
	return m.name
}

// GetValue 値を返す
func (m defaultNode) GetValue() string {
	return m.value
}

// GetChildren 子ノードを返す
func (m defaultNode) GetChildren() []Node {
	return m.children
}

// GetStyle Presentationを返す
func (m defaultNode) GetStyle() MarkupStyle {
	return m.style
}
