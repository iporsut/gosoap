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

type SOAPBody struct {
	Use string `xml:"use,attr"`
}

type Operation struct {
	Message string `xml:"message,attr"`
}

type IOOperation struct {
	Operation
	Body SOAPBody `xml:"body"`
}

type InputOperation struct {
	IOOperation
}

type OutputOperation struct {
	IOOperation
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
	MinOccurs int `xml:"minOccurs"`
	MaxOccurs int `xml:"maxOccurs"`
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

type Address struct {
	Location string `xml:"location,attr"`
}

type Port struct {
	Binding string  `xml:"binding,attr"`
	Name    string  `xml:"name,attr"`
	Address Address `xml:"address"`
}

type Service struct {
	Name string `xml:"name,attr"`
	Port Port   `xml:"port"`
}

type SOAPBinding struct {
	Style     string `xml:"style,attr"`
	Transport string `xml:"transport,attr"`
}

type Binding struct {
	Name      string        `xml:"name,attr"`
	Type      string        `xml:"type,attr"`
	Binding   SOAPBinding   `xml:"binding"`
	Operation WSDLOperation `xml:"operation"`
}

type Definition struct {
	XMLName       xml.Name  `xml:"definitions"`
	Documentation string    `xml:"documentation"`
	Messages      []Message `xml:"message"`
	PortType      PortType  `xml:"portType"`
	Types         Types     `xml:"types"`
	Service       Service   `xml:"service"`
	Binding       Binding   `xml:"binding"`
}

func Unmarshal(b []byte) (definition Definition, err error) {
	err = xml.Unmarshal(b, &definition)
	return
}
