package converter

import (
	"encoding/xml"
	"strings"
)

// Node XMLを表現する
type Node struct {
	Name  xml.Name
	Attr  []xml.Attr
	Value string
	Nodes []*Node
}

// UnmarshalXML XMLからデコード
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
	e.Name = start.Name
	e.Attr = start.Attr
	e.Nodes = nodes
	return nil
}

// MarshalXML XMLにエンコード
func (e *Node) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name = e.Name
	start.Attr = e.Attr
	return enc.EncodeElement(struct {
		Data  string `xml:",chardata"`
		Nodes []*Node
	}{
		Data:  e.Value,
		Nodes: e.Nodes,
	}, start)
}
