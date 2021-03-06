package mathml

import "strings"

// StringWithNoAttr Nodeを文字列として出力します
func StringWithNoAttr(node *Node) string {

	builder := strings.Builder{}
	stringWithNoAttr(node, &builder)
	return builder.String()
}

func stringWithNoAttr(node *Node, builder *strings.Builder) {

	builder.WriteString("<")
	builder.WriteString(node.Name)
	builder.WriteString(">")
	builder.WriteString(node.Value)
	for _, child := range node.Children {
		stringWithNoAttr(child, builder)
	}
	builder.WriteString("</")
	builder.WriteString(node.Name)
	builder.WriteString(">")
}

// StringWithAttr 属性を含めて出力します
func StringWithAttr(node *Node) string {

	builder := strings.Builder{}
	stringWithAttr(node, &builder)
	return builder.String()
}

func stringWithAttr(node *Node, builder *strings.Builder) {
	builder.WriteString("<")
	builder.WriteString(node.Name)
	for _, attr := range node.Attrs {
		builder.WriteString(" ")
		builder.WriteString(attr.Key)
		builder.WriteString("=\"")
		builder.WriteString(attr.Value)
		builder.WriteString("\"")
	}
	builder.WriteString(">")
	for _, child := range node.Children {
		stringWithAttr(child, builder)
	}
	builder.WriteString(node.Value)
	builder.WriteString("</")
	builder.WriteString(node.Name)
	builder.WriteString(">")
}
