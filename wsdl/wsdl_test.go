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

func TestUnmarshalMessage(t *testing.T) {
	var b, err = ioutil.ReadFile("./testdata/wsdl.xml")
	if err != nil {
		t.Error(err)
	}

	definition, err := Unmarshal(b)
	if err != nil {
		t.Error(err)
	}

	messages := definition.Messages

	if messages[0].Name != "ReadRetlWSInput" {
		t.Errorf("expect \"ReadRetlWSInput\" but was %s", messages[0].Name)
	}
	if messages[1].Name != "ReadRetlWSOutput" {
		t.Errorf("expect \"ReadRetlWSOutput\" but was %s", messages[0].Name)
	}
	if messages[2].Name != "ReadRetlWSError" {
		t.Errorf("expect \"ReadRetlWSError\" but was %s", messages[0].Name)
	}
}
