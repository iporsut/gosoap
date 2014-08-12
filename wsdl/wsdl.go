package wsdl

import (
	"encoding/xml"
)

type Part struct {
	Element string `xml:"element,attr"`
	Name    string `xml:"name,attr"`
}

type Message struct {
	Name string `xml:"name,attr"`
	Part Part   `xml:"part"`
}

type Operation struct {
	Message string `xml:"message,attr"`
}

type InputOperation struct {
	Operation
}

type OutputOperation struct {
	Operation
}

type FaultOperation struct {
	Operation
	Name string `xml:"name,attr"`
}

type WSDLOperation struct {
	Name   string          `xml:"name,attr"`
	Input  InputOperation  `xml:"input"`
	Output OutputOperation `xml:"output"`
	Fault  FaultOperation  `xml:"fault"`
}

type PortType struct {
	Name       string          `xml:"name,attr"`
	Operations []WSDLOperation `xml:"operation"`
}

type Element struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type SequenceElement struct {
	Element
}

type Sequence struct {
	Elements []SequenceElement `xml:"element"`
}

type ComplexType struct {
	Name     string   `xml:"name,attr"`
	Sequence Sequence `xml:"sequence"`
}

type SchemaElement struct {
	Element
	ComplexType ComplexType `xml:"complexType"`
}

type Schema struct {
	AttributeFormDefault string          `xml:"attributeFormDefault,attr"`
	ElementFormDefault   string          `xml:"elementFormDefault,attr"`
	TargetNamespace      string          `xml:"targetNamespace,attr"`
	Elements             []SchemaElement `xml:"element"`
	ComplexTypes         []ComplexType   `xml:"complexType"`
}

type Types struct {
	Schema Schema `xml:"schema"`
}

type Definition struct {
	XMLName       xml.Name  `xml:"definitions"`
	Documentation string    `xml:"documentation"`
	Messages      []Message `xml:"message"`
	PortType      PortType  `xml:"portType"`
	Types         Types     `xml:"types"`
}

func Unmarshal(b []byte) (definition Definition, err error) {
	err = xml.Unmarshal(b, &definition)
	return
}
