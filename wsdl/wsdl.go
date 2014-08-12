package wsdl

import (
	"encoding/xml"
)

type Part struct {
	Element string `xml:"element,attr"`
	Name    string `xml:"name,attr"`
}

type Message struct {
	XMLName xml.Name `xml:"message"`
	Name    string   `xml:"name,attr"`
	Part    Part     `xml:"part"`
}

type Definition struct {
	XMLName       xml.Name  `xml:"definitions"`
	Documentation string    `xml:"documentation"`
	Messages      []Message `xml:"message"`
}

func Unmarshal(b []byte) (definition Definition, err error) {
	err = xml.Unmarshal(b, &definition)
	return
}
