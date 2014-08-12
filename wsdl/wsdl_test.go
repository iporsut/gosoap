package wsdl

import (
	"io/ioutil"
	"testing"
)

func TestUnmarshalDocumentation(t *testing.T) {
	var b, err = ioutil.ReadFile("./testdata/wsdl.xml")
	if err != nil {
		t.Error(err)
	}

	definition, err := Unmarshal(b)

	if definition.Documentation != "Generated at 08-11-2014 16:33:06:253" {
		t.Errorf("expect \"Generated at 08-11-2014 16:33:06:253\" but was %s", definition.Documentation)
	}
}
