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

type InputOperation struct {
	Message string `xml:"message,attr"`
}

type OutputOperation struct {
	Message string `xml:"message,attr"`
}

type FaultOperation struct {
	Message string `xml:"message,attr"`
	Name    string `xml:"name,attr"`
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
}

type Schema struct {
	AttributeFormDefault string    `xml:"attributeFormDefault,attr"`
	ElementFormDefault   string    `xml:"elementFormDefault,attr"`
	TargetNamespace      string    `xml:"targetNamespace,attr"`
	Elements             []Element `xml:"element"`
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
