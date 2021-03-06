package mathml

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
	Attrs    []*Attr
	Children []*Node
	Style    MarkupStyle
}

// Attr MathMLにおける属性を表す
type Attr struct {
	Key   string
	Value string
}

// Map 下にある全てのノードについてcallbackを実行し結果をsliceで返す
func (n *Node) Map(callback func(*Node) interface{}) []interface{} {
	slice := make([]interface{}, 0)
	n.ForEach(func(n *Node) {
		slice = append(slice, callback(n))
	})
	return slice
}

// List を取得します
func (n *Node) List() []*Node {
	var nodes []*Node
	n.ForEach(func(n *Node) {
		nodes = append(nodes, n)
	})
	return nodes
}

// ForEach 深さ優先で探索します
func (n *Node) ForEach(callback func(*Node)) {
	callback(n)
	for _, child := range n.Children {
		child.ForEach(callback)
	}
}
