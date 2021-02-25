package converter

import "prototype.mathbase.app/mathml"

func mathMLNodeFactory(node *xmlNode) *mathml.Node {

	children := []*mathml.Node{}
	for _, n := range node.Nodes {
		children = append(children, mathMLNodeFactory(n))
	}

	mNode := mathml.Node{Name: node.Name, Value: node.Value, Children: children, Style: mathml.Presentation}
	return &mNode
}
