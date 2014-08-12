package wsdl

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func testRead(t *testing.T, filename string) []byte {
	var b, err = ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return []byte{}
	}
	return b
}

func testUnmarshal(t *testing.T, b []byte) Definition {
	definition, err := Unmarshal(b)
	if err != nil {
		t.Error(err)
		return Definition{}
	}
	return definition
}

func testUnmarshalFromFile(t *testing.T, filename string) (definition Definition) {
	b := testRead(t, "./testdata/wsdl.xml")
	definition = testUnmarshal(t, b)
	return
}

type testCase func(t *testing.T, definition Definition)

func (test testCase) call(t *testing.T, definition Definition) {
	test(t, definition)
}

func TestUnmarshal(t *testing.T) {
	var definition Definition = testUnmarshalFromFile(t, "./testdata/wsdl.xml")

	var tests = []testCase{
		testCase(testUnmarshalDocumentation),
		testCase(testUnmarshalMessage),
		testCase(testUnmarshalPartInMessage),
		testCase(testUnmarshalPortType),
		testCase(testUnmarshalOperationInPortType),
		testCase(testUnmarshalOperationInputInPortType),
		testCase(testUnmarshalOperationOutputInPortType),
		testCase(testUnmarshalOperationFault),
		testCase(testUnmarshalSchema),
		testCase(testUnmarshalElement),
		testCase(testUnmarshalComplexTypeInSchema),
		testCase(testUnmarshalService),
		testCase(testUnmarshalPortInService),
	}

	for _, test := range tests {
		test.call(t, definition)
	}
}

func testUnmarshalDocumentation(t *testing.T, definition Definition) {
	if definition.Documentation != "Generated at 08-11-2014 16:33:06:253" {
		t.Errorf("expect \"Generated at 08-11-2014 16:33:06:253\" but was %s", definition.Documentation)
	}
}

func testUnmarshalMessage(t *testing.T, definition Definition) {
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

func testUnmarshalPartInMessage(t *testing.T, definition Definition) {
	messages := definition.Messages
	parts := []Part{
		messages[0].Part,
		messages[1].Part,
		messages[2].Part,
	}

	expectedParts := []Part{
		Part{
			Element: "tuxtype:ReadRetlWS",
			Name:    "FML32",
		},
		Part{
			Element: "tuxtype:ReadRetlWSResponse",
			Name:    "FML32",
		},
		Part{
			Element: "tuxtype:ReadRetlWSFault",
			Name:    "FML32",
		},
	}

	if !reflect.DeepEqual(parts, expectedParts) {
		t.Errorf("expect parts %v but was %v", expectedParts, parts)
	}
}

func testUnmarshalPortType(t *testing.T, definition Definition) {
	portType := definition.PortType

	if portType.Name != "INReadRetlWS_PortType" {
		t.Errorf("expect \"INReadRetlWS_PortType\" but was %s", portType.Name)
	}
}

func testUnmarshalOperationInPortType(t *testing.T, definition Definition) {
	var operations []WSDLOperation = definition.PortType.Operations

	if operations[0].Name != "ReadRetlWS" {
		t.Errorf("expect \"ReadRetlWS\" but was %s", operations[0].Name)
	}
}

func testUnmarshalOperationInputInPortType(t *testing.T, definition Definition) {
	var inputOperation InputOperation = definition.PortType.Operations[0].Input

	if inputOperation.Message != "tns:ReadRetlWSInput" {
		t.Errorf("expect \"tns:ReadRetlWSInput\" but was %s", inputOperation.Message)
	}
}

func testUnmarshalOperationOutputInPortType(t *testing.T, definition Definition) {
	var outputOperation OutputOperation = definition.PortType.Operations[0].Output

	if outputOperation.Message != "tns:ReadRetlWSOutput" {
		t.Errorf("expect \"tns:ReadRetlWSOutput\" but was %s", outputOperation.Message)
	}
}

func testUnmarshalOperationFault(t *testing.T, definition Definition) {
	var faultOperation FaultOperation = definition.PortType.Operations[0].Fault

	if faultOperation.Message != "tns:ReadRetlWSError" {
		t.Errorf("expect \"tns:ReadRetlWSError\" but was %s", faultOperation.Message)
	}

	if faultOperation.Name != "ReadRetlWSError" {
		t.Errorf("expect \"ReadRetlWSError\" but was %s", faultOperation.Name)
	}
}

func testUnmarshalSchema(t *testing.T, definition Definition) {
	var schema Schema = definition.Types.Schema

	if schema.AttributeFormDefault != "unqualified" {
		t.Errorf("expect \"unqualified\" but was %s", schema.AttributeFormDefault)
	}

	if schema.ElementFormDefault != "qualified" {
		t.Errorf("expect \"unqualified\" but was %s", schema.ElementFormDefault)
	}

	if schema.TargetNamespace != "urn:pack.INReadRetlWS_typedef.salt11" {
		t.Errorf("expect \"urn:pack.INReadRetlWS_PortType.salt11\" but was %s", schema.TargetNamespace)
	}
}

func testUnmarshalElement(t *testing.T, definition Definition) {
	var elements []SchemaElement = definition.Types.Schema.Elements

	if len(elements) != 3 {
		t.Errorf("expect 3 but was %d", len(elements))
	}

	var names = []string{"ReadRetlWS", "ReadRetlWSResponse", "ReadRetlWSFault"}
	for i := 0; i < len(names); i++ {
		if elements[i].Name != names[i] {
			t.Errorf("expect %s but was %s", names[i], elements[i].Name)
		}
	}
}

func testUnmarshalComplexTypeInElement(t *testing.T, definition Definition) {
	var element SequenceElement = definition.Types.Schema.Elements[0].ComplexType.Sequence.Elements[0]

	if element.Name != "inbuf" {
		t.Errorf("expect \"inbuf\" but was %s", element.Name)
	}

	if element.Type != "tuxtype:fml32_ReadRetlWS_In" {
		t.Errorf("expect \"tuxtype:fml32_ReadRetlWS_In\" but was %s", element.Type)
	}
}

func testUnmarshalComplexTypeInSchema(t *testing.T, definition Definition) {
	var complexTypes []ComplexType = definition.Types.Schema.ComplexTypes

	var names = []string{"fml32_ReadRetlWS_In", "fml32_ReadRetlWS_Out", "fml32_ReadRetlWS_Err"}
	for i := 0; i < len(names); i++ {
		if complexTypes[i].Name != names[i] {
			t.Errorf("expect %s but was %s", names[i], complexTypes[i].Name)
		}
	}
}

func testUnmarshalAttributeElementInComplexTypeSchema(t *testing.T, definition Definition) {
	var elements []SequenceElement = definition.Types.Schema.ComplexTypes[0].Sequence.Elements

	var expectedElements = []SequenceElement{
		SequenceElement{
			Element: Element{
				Name: "USER_CODE",
				Type: "xsd:string",
			},
			MinOccurs: 0,
			MaxOccurs: 1,
		},
		SequenceElement{
			Element: Element{
				Name: "RD_RETL__RETL_CODE",
				Type: "xsd:string",
			},
			MinOccurs: 0,
			MaxOccurs: 1,
		},
		SequenceElement{
			Element: Element{
				Name: "RD_RETL__RETL_NAME",
				Type: "xsd:string",
			},
			MinOccurs: 0,
			MaxOccurs: 1,
		},
		SequenceElement{
			Element: Element{
				Name: "RD_RETL__SHOP_NAME",
				Type: "xsd:string",
			},
			MinOccurs: 0,
			MaxOccurs: 1,
		},
		SequenceElement{
			Element: Element{
				Name: "RD_RETL__RETL_STTS",
				Type: "xsd:string",
			},
			MinOccurs: 0,
			MaxOccurs: 1,
		},
		SequenceElement{
			Element: Element{
				Name: "READ_FLAG",
				Type: "xsd:string",
			},
			MinOccurs: 0,
			MaxOccurs: 1,
		},
	}

	if !reflect.DeepEqual(elements, expectedElements) {
		t.Errorf("expect %v but was %v", expectedElements, elements)
	}
}

func testUnmarshalService(t *testing.T, definition Definition) {
	var service Service = definition.Service

	if service.Name != "TuxedoWebService" {
		t.Errorf("expect \"TuxedoWebService\" but was %s", service.Name)
	}
}

func testUnmarshalPortInService(t *testing.T, definition Definition) {
	var port = definition.Service.Port

	if port.Binding != "tns:INReadRetlWS_Binding" {
		t.Errorf("expect \"tns:INReadRetlWS_Binding\" but was %s", port.Binding)
	}

	if port.Name != "INReadRetlWS_Endpoint" {
		t.Errorf("expect \"INReadRetlWS_Endpoint\" but was %s", port.Name)
	}
}
