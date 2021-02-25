package converter

import "prototype.mathbase.app/mathml"

const mathmlNs = "http://www.w3.org/1998/Math/MathML"

func mathMLFactory(node *xmlNode) *mathml.Node {

	if node.Name.Space != mathmlNs {
		panic("parsed document may not have MathML namespace.")
	}

	children := []*mathml.Node{}
	for _, n := range node.Nodes {
		children = append(children, mathMLFactory(n))
	}

	mNode := mathml.Node{Name: node.Name.Local, Value: node.Value, Children: children, Style: mathml.Presentation}
	return &mNode
}
