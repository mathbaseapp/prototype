package converter

import (
	"errors"

	"prototype.mathbase.app/mathml"
)

const mathmlNs = "http://www.w3.org/1998/Math/MathML"

func mathMLFactory(node *xmlNode) (*mathml.Node, error) {

	if node.Name.Space != mathmlNs {
		return nil, errors.New("parsed document may not have MathML namespace")
	}

	children := []*mathml.Node{}
	for _, n := range node.Nodes {
		mm, err := mathMLFactory(n)
		if err != nil {
			return nil, err
		}
		children = append(children, mm)
	}
	attrs := []*mathml.Attr{}
	for _, attr := range node.Attr {
		attrs = append(attrs, &mathml.Attr{Key: attr.Name.Local, Value: attr.Value})
	}
	mNode := mathml.Node{Name: node.Name.Local, Value: node.Value, Children: children, Style: mathml.Presentation, Attrs: attrs}
	return &mNode, nil
}
