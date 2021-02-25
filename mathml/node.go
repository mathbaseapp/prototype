package mathml

import (
	"encoding/xml"
)

// MarkupStyle マークアップの表現方法を示します
type MarkupStyle string

const (
	// Presentation プレゼンテーションマークアップ
	Presentation = MarkupStyle("Presentation")
	// Content コンテンツマークアップ
	Content = MarkupStyle("Content")
)

// Node MathMLにおけるNodeを表現します
type Node struct {
	Name     xml.Name
	Value    string
	Children []*Node
	Style    MarkupStyle
}
