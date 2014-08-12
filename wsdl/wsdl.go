package wsdl

import (
	"encoding/xml"
)

type Definition struct {
	XMLName       xml.Name `xml:"definitions"`
	Documentation string   `xml:"documentation"`
}

func Unmarshal(b []byte) (definition Definition, err error) {
	err = xml.Unmarshal(b, &definition)
	return
}
