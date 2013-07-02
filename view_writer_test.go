package main

import (
	"log"
	// "os"
	"bytes"
	"fmt"
	"testing"
)

func Test_ViewWriterHookup(t *testing.T) {
	log.Printf("Hookup Succeeded - testing view writer")
}

func Test_SmokeTest(t *testing.T) {
	parser := NewParser("Writer smoke test", ".cl1.cl2#id1 Some text\n\t%p Child paragraph")
	parser.Parse()

	//vw := newViewWriter(os.Stdout, parser.root, "Test")
	//vw.WriteView()
}

func Test_GetCharStr15(t *testing.T) {
	str := genString(15, 3)
	if str != "12 45 78 01 34 " {
		t.Error("String not as expected: " + str)
	}

	str = genString(20, 10)
	if str != "123456789 123456789 " {
		t.Error("String not as expected: " + str)
	}
}

func Test_GetWhitespaceIndicies(t *testing.T) {
	str := genString(15, 3)

	vw := NewViewWriter(nil, nil, nil, "")
	spaces := vw.getWhitespaceIndicies(str)

	if len(spaces) != 5 {
		t.Error("Number of spaces not as expected: %d", spaces)
	}

	chkSpace(0, 2, spaces, t)
	chkSpace(1, 5, spaces, t)
	chkSpace(2, 8, spaces, t)
	chkSpace(3, 11, spaces, t)
	chkSpace(4, 14, spaces, t)
}

func chkSpace(sliceIndex int, expectedValue int, spaces []int, t *testing.T) {
	if spaces[sliceIndex] != expectedValue {
		t.Error(fmt.Sprintf("Unexpected value at index %d: %d", sliceIndex, spaces[sliceIndex]))
	}
}

func Test_FormattingShortText(t *testing.T) {
	txt := genString(9, 10)
	expected := `123456789
`
	testFormatStr(txt, expected, t)
}

func Test_FormattingOneWrap(t *testing.T) {
	txt := genString(130, 10)
	expected :=
		`123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 
123456789 
`

	testFormatStr(txt, expected, t)
}

func genString(count int, moduloToSpace int) string {
	buffer := bytes.NewBuffer(make([]byte, 0))
	for i := 1; i <= count; i++ {
		if i%moduloToSpace == 0 {
			fmt.Fprintf(buffer, " ")
		} else {
			fmt.Fprintf(buffer, "%d", i%10)
		}
	}
	return buffer.String()
}

func testFormatStr(s string, expected string, t *testing.T) {
	vw := NewViewWriter(nil, nil, nil, "")
	buffer := bytes.NewBuffer(make([]byte, 0))
	ind := NewIndentingWriter(buffer)

	vw.writeLongText(s, ind)

	result := buffer.String()
	if result != expected {
		t.Error("result not as expected")
		t.Log(fmt.Sprintf("Expected: [%s]", expected))
		t.Log(fmt.Sprintf("Output: [%s]", result))
	}
}
