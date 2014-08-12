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

type Definition struct {
	XMLName       xml.Name  `xml:"definitions"`
	Documentation string    `xml:"documentation"`
	Messages      []Message `xml:"message"`
	PortType      PortType  `xml:"portType"`
}

func Unmarshal(b []byte) (definition Definition, err error) {
	err = xml.Unmarshal(b, &definition)
	return
}
