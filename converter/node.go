package converter

import (
	"encoding/xml"
	"strings"
)

type xmlNode struct {
	Name  xml.Name
	Attr  []xml.Attr
	Value string
	Nodes []*xmlNode
}

// UnmarshalXML XMLからデコード
func (e *xmlNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var nodes []*xmlNode
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
			e := &xmlNode{}
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
func (e *xmlNode) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name = e.Name
	start.Attr = e.Attr
	return enc.EncodeElement(struct {
		Data  string `xml:",chardata"`
		Nodes []*xmlNode
	}{
		Data:  e.Value,
		Nodes: e.Nodes,
	}, start)
}
