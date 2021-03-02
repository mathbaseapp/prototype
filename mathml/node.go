package mathml

import "prototype.mathbase.app/middlelng"

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
	Name     string
	Value    string
	Children []*Node
	Style    MarkupStyle
}

// Map 下にある全てのノードについてcallbackを実行し結果をsliceで返す
func (n *Node) Map(callback func(middlelng.MiddleLanguage) interface{}) []interface{} {
	slice := make([]interface{}, 0)
	slice = append(slice, callback(n))
	for _, child := range n.Children {
		slice = append(slice, child.Map(callback)...)
	}
	return slice
}
