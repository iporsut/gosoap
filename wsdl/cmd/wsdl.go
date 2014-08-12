package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	//"net/http"
)

var namespaces map[string]string = make(map[string]string)

type Sequence struct {
	Elements []Element `xml:"element"`
}

type ComplexType struct {
	Name     string   `xml:"name,attr"`
	Sequence Sequence `xml:"sequence"`
}

type Element struct {
	Name        string      `xml:"name,attr"`
	MaxOccurs   string      `xml:"maxOccurs,attr"`
	MinOccurs   string      `xml:"minOccurs,attr"`
	Type        string      `xml:"type,attr"`
	Ref         string      `xml:"ref,attr"`
	ComplexType ComplexType `xml:"complexType"`
}

type Schema struct {
	AttributeFormDefault string        `xml:"attributeFormDefault,attr"`
	ElementFormDefault   string        `xml:"elementFormDefault,attr"`
	TargetNamespace      string        `xml:"targetNamespace,attr"`
	Elements             []Element     `xml:"element"`
	ComplexType          []ComplexType `xml:"complexType"`
}

type Types struct {
	Schemas []Schema `xml:"schema"`
}

type Part struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type Message struct {
	Name string `xml:"name,attr"`
	Part Part   `xml:"part"`
}

type SOAPBody struct {
	Use string `xml:"use,attr"`
}

type WSDLInput struct {
	Message string   `xml:"message,attr"`
	Name    string   `xml:"name,attr"`
	Body    SOAPBody `xml:"body"`
}

type WSDLOutput struct {
	Message string   `xml:"message,attr"`
	Name    string   `xml:"name,attr"`
	Body    SOAPBody `xml:"body"`
}

type WSDLFault struct {
	Message string    `xml:"message,attr"`
	Name    string    `xml:"name,attr"`
	Fault   SOAPFault `xml:"fault"`
}

type SOAPFault struct {
	Name string `xml:"name,attr"`
	Use  string `xml:"use,attr"`
}

type SOAPOperation struct {
	SOAPAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

type WSDLOperation struct {
	Name      string        `xml:"name,attr"`
	Input     WSDLInput     `xml:"input"`
	Output    WSDLOutput    `xml:"output"`
	Fault     WSDLFault     `xml:"fault"`
	Operation SOAPOperation `xml:"operation"`
}

type PortType struct {
	Name       string          `xml:"name,attr"`
	Operations []WSDLOperation `xml:"operation"`
}

type SOAPBinding struct {
	Style     string `xml:"style,attr"`
	Transport string `xml:"transport,attr"`
}

type WSDLBinding struct {
	Name       string          `xml:"name,attr"`
	Type       string          `xml:"type,attr"`
	Binding    SOAPBinding     `xml:"binding"`
	Operations []WSDLOperation `xml:"operation"`
}

type Address struct {
	Location string `xml:"location,attr"`
}

type Port struct {
	Name    string  `xml:"name,attr"`
	Binding string  `xml:"binding,attr"`
	Address Address `xml:"address"`
}

type Service struct {
	Name string `xml:"name,attr"`
	Port Port   `xml:"port"`
}

type WSDL struct {
	Documentaion string      `xml:"documentation"`
	Messages     []Message   `xml:"message"`
	PortType     PortType    `xml:"portType"`
	Service      Service     `xml:"service"`
	Binding      WSDLBinding `xml:"binding"`
	Types        Types       `xml:"types"`
}

func main() {
	/*resp, err := http.Get("http://www.herongyang.com/Service/Hello_WSDL_11_SOAP.wsdl")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)*/
	//b, err := ioutil.ReadFile("NCCAQueryAssetNote.wsdl")
	b, err := ioutil.ReadFile("wsdl.xml")
	if err != nil {
		log.Fatal(err)
	}
	wsdl := new(WSDL)
	err = xml.Unmarshal(b, wsdl)
	if err != nil {
		log.Fatal(err)
	}

	for _, op := range wsdl.PortType.Operations {
		inputMessageName := strings.Split(op.Input.Message, ":")[1]
		inputMessage := wsdl.FindMessageByName(inputMessageName)
		elementName := strings.Split(inputMessage.Part.Element, ":")[1]
		element, schema := wsdl.FindElementByName(elementName)
		namespace := addNamespace(schema.TargetNamespace)
		wsdl.PrintElementType(element, namespace, os.Stdout)
	}
}

func addNamespace(namespace string) string {
	paths := strings.Split(namespace[:len(namespace)-1], "/")
	lastPath := paths[len(paths)-1]
	prefixNumber := 3
	prefix := lastPath[:prefixNumber]
	for old, ok := namespaces[prefix]; ok && (old != namespace); old, ok = namespaces[prefix] {
		prefixNumber++
		prefix = lastPath[:prefixNumber]
	}
	namespaces[prefix] = namespace
	return prefix
}

func (w *WSDL) FindMessageByName(name string) *Message {
	for _, m := range w.Messages {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (w *WSDL) FindElementByName(name string) (*Element, *Schema) {
	for _, s := range w.Types.Schemas {
		for _, e := range s.Elements {
			if e.Name == name {
				return &e, &s
			}
		}
	}
	return nil, nil
}

func (w *WSDL) FindComplexTypeByName(name string) (*ComplexType, *Schema) {
	for _, s := range w.Types.Schemas {
		for _, c := range s.ComplexType {
			if c.Name == name {
				return &c, &s
			}
		}
	}
	return nil, nil
}

func (w *WSDL) PrintElementType(e *Element, namespace string, wt io.Writer) {
	if e.Ref != "" && e.Name == "" {
		refName := strings.Split(e.Ref, ":")[1]
		e, schema := w.FindElementByName(refName)
		newNameSpace := addNamespace(schema.TargetNamespace)
		w.PrintElementType(e, newNameSpace, wt)
	} else {
		fmt.Fprintf(wt, "<%s:%s>", namespace, e.Name)
		if len(e.ComplexType.Sequence.Elements) > 0 {
			fmt.Fprint(wt, "\n")
			for _, se := range e.ComplexType.Sequence.Elements {
				w.PrintElementType(&se, namespace, wt)
			}
		}

		if e.Type != "" {
			typeName := strings.Split(e.Type, ":")[1]
			c, _ := w.FindComplexTypeByName(typeName)

			if c == nil {
				fmt.Fprint(wt, "?")
			} else {
				fmt.Fprint(wt, "\n")
				for _, se := range c.Sequence.Elements {
					w.PrintElementType(&se, namespace, wt)
				}
			}
		}
		fmt.Fprintf(wt, "</%s:%s>\n", namespace, e.Name)
	}
}
