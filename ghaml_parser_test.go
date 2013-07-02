package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_ParserHookup(t *testing.T) {
	log.Printf("Hookup Succeeded - testing parsing")
}

func getTestParser(testname, content string) *GhamlParser {
	return NewParser(testname, content)
}

func Test_ParserSetup(t *testing.T) {
	parser := getTestParser("Test parser", "%test")

	if parser.root == nil {
		t.Error("Root node is null")
	}

	if parser.tags == nil {
		t.Error("tags stack is null")
	}

	if parser.tags.count() != 1 {
		t.Error("Tags stack is not as expected")
	}
}

func Test_SimpleParser(t *testing.T) {
	hamlStr := `%test`
	expStr := `test`

	parseAndCompare(hamlStr, expStr, t)
}

func Test_MultipleTags(t *testing.T) {
	hamlStr := `%tag1
%tag2`
	expStr := `tag1
tag2`

	parseAndCompare(hamlStr, expStr, t)
}

func Test_ChildTag(t *testing.T) {
	hamlStr := `%tag1
    %tag2`

	expStr := "tag1\n\ttag2"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_Nesting(t *testing.T) {
	hamlStr := "%t_1\n\t%t_1_1\n\t\t%t_1_1_1\n\t%t_1_2\n%t_2\n\t%t_2_1"
	expStr := "t_1\n\tt_1_1\n\t\tt_1_1_1\n\tt_1_2\nt_2\n\tt_2_1"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_TagWithText(t *testing.T) {
	hamlStr := "%p a paragraph\n%p second paragraph"
	expStr := "p (a paragraph)\np (second paragraph)"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_TagWithIndentedText(t *testing.T) {
	hamlStr := "%p a paragraph\n\twith indented text"
	expStr := "p (a paragraph with indented text)"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_ParserDoctype(t *testing.T) {
	hamlStr := "!!!"
	expStr := "doctype"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_TagWithClass(t *testing.T) {
	hamlStr := "%tag.testClass"
	expStr := "tag class='testClass'"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_TagWithId(t *testing.T) {
	hamlStr := "%tag#testId"
	expStr := "tag id='testId'"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_ParserTagWithClassAndId(t *testing.T) {
	hamlStr := "%tag#testId.testClass"
	expStr := "tag id='testId' class='testClass'"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_TagWithMultipleClasses(t *testing.T) {
	hamlStr := "%tag1.class1.class2"
	expStr := "tag1 class='class1, class2'"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_ParserImplicitDiv(t *testing.T) {
	hamlStr := "#testId\n.testClass"
	expStr := "div id='testId'\ndiv class='testClass'"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_ImplicitWithMultipleClasses(t *testing.T) {
	hamlStr := ".cl1.cl2#id1 Some text"
	expStr := "div id='id1' class='cl1, cl2' (Some text)"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_Attributes(t *testing.T) {
	hamlStr := "%tag1{\"class\":\"yes\", width: 5px} content"
	expStr := "tag1 class='yes' width='5px' (content)"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_CodeOutput(t *testing.T) {
	hamlStr := "= fmt.Println(\"hi\")"
	expStr := "code_output (fmt.Println(\"hi\"))"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_TagAndCodeOutput(t *testing.T) {
	hamlStr := "%p.hi= fmt.Println(\"hi\")"
	expStr := "p class='hi'\n\tcode_output (fmt.Println(\"hi\"))"

	parseAndCompare(hamlStr, expStr, t)
}

func Test_ParserDataType(t *testing.T) {
	hamlStr := "@data_type: []string"
	parser := NewParser(hamlStr, hamlStr)
	parser.Parse()

	if bytes.Compare([]byte(parser.context.dataType), []byte("[]string")) != 0 {
		t.Error("datatype parse error. Expected '[]string', got: " + parser.context.dataType)
	}
}

func Test_ParserImport(t *testing.T) {
	hamlStr := "@import ( \"fmt\" ) "
	parser := NewParser(hamlStr, hamlStr)
	parser.Parse()

	if bytes.Compare([]byte(parser.context.imports[0]), []byte("fmt")) != 0 {
		t.Error("datatype parse error. Expected '[]string', got: " + parser.context.dataType)
	}
}

func Test_ParserMultipleImports(t *testing.T) {
	hamlStr := "@import (\n\t\"fmt\"\n\t\"strings\") "
	parser := NewParser(hamlStr, hamlStr)
	parser.Parse()

	if bytes.Compare([]byte(parser.context.imports[0]), []byte("fmt")) != 0 {
		t.Error("datatype parse error. Expected '[]string', got: " + parser.context.dataType)
	}

	if bytes.Compare([]byte(parser.context.imports[1]), []byte("strings")) != 0 {
		t.Error("datatype parse error. Expected '[]string', got: " + parser.context.dataType)
	}
}

func Test_Metadata(t *testing.T) {
	hamlStr := `
@import (
  "fmt"
  "os"
)

@data_type: test_struct

%html
  %head
    %title= test_struct.title
  %body
    .main= test_struct.body
`

	parser := NewParser(hamlStr, hamlStr)
	parser.Parse()

}

func dumpParser(parser *GhamlParser) {
	res := parser.dumpNodes()
	log.Printf("Dumping parser %q:\n%s", parser.name, res)
}

func parseAndCompare(markup, expectedDump string, t *testing.T) {
	parser := NewParser(markup, markup)
	parser.Parse()

	dumpStr := parser.dumpNodes()
	if bytes.Compare([]byte(expectedDump), []byte(dumpStr)) != 0 {
		t.Error("expected: [" + expectedDump + "], received: [" + dumpStr + "]")
	}
}
