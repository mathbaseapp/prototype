package mathml

import "strings"

// Printer Nodeを文字列として出力します
func Printer(node Node) string {

	builder := strings.Builder{}
	printer(node, &builder)
	return builder.String()
}

func printer(node Node, builder *strings.Builder) {

	builder.WriteString("<")
	builder.WriteString(node.GetName().Local)
	builder.WriteString(">")
	builder.WriteString(node.GetValue())
	for _, child := range node.GetChildren() {
		printer(child, builder)
	}
	builder.WriteString("</")
	builder.WriteString(node.GetName().Local)
	builder.WriteString(">")
}
