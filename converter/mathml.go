package converter

import (
	"errors"

	"prototype.mathbase.app/mathml"
)

const mathmlNs = "http://www.w3.org/1998/Math/MathML"

func mathMLFactory(node *xmlNode) (*mathml.Node, error) {

	if node.Name.Space != mathmlNs {
		return &mathml.Node{}, errors.New("parsed document may not have MathML namespace")
	}

	children := []*mathml.Node{}
	for _, n := range node.Nodes {
		mm, _ := mathMLFactory(n)
		children = append(children, mm)
	}

	mNode := mathml.Node{Name: node.Name.Local, Value: node.Value, Children: children, Style: mathml.Presentation}
	return &mNode, nil
}
