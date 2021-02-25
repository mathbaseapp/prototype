package converter

import "prototype.mathbase.app/mathml"

func defaultNodeFactory(node *xmlNode) mathml.Node {

	children := []mathml.Node{}
	for _, n := range node.Nodes {
		children = append(children, defaultNodeFactory(n))
	}

	return mathml.DefaultNode(node.Name, node.Value, children, mathml.Presentation)
}
