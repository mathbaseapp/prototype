package main

import (
	"encoding/xml"
	"fmt"
	"os/exec"
	"strings"
)

type Node struct {
	XMLName    xml.Name
	Attributes []xml.Attr
	Value      string
	Nodes      []*Node
}

func (e *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var nodes []*Node
	var done bool
	for !done {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case xml.CharData:
			e.Value = strings.TrimSpace(string(t))
		case xml.StartElement:
			e := &Node{}
			e.UnmarshalXML(d, t)
			nodes = append(nodes, e)
		case xml.EndElement:
			done = true
		}
	}
	e.XMLName = start.Name
	e.Attributes = start.Attr
	e.Nodes = nodes
	return nil
}

func (e *Node) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name = e.XMLName
	start.Attr = e.Attributes
	return enc.EncodeElement(struct {
		Data  string `xml:",chardata"`
		Nodes []*Node
	}{
		Data:  e.Value,
		Nodes: e.Nodes,
	}, start)
}

func main() {
	latexStr := "$$x = 4$$"
	pandocCmd := "echo '" + latexStr + "'  | pandoc -f html+tex_math_dollars -t html --mathml"
	out, err := exec.Command("sh", "-c", pandocCmd).Output()
	str := Node{}
	fmt.Println(string(out))
	xml.Unmarshal(out, &str)

	if err != nil {
		fmt.Println("Command Exec Error.")
	}
	fmt.Println(str.Nodes[0].Nodes[0].Nodes[0].Value)
}
