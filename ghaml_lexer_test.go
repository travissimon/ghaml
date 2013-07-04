package main

import (
	"log"
	"testing"
)

func Test_LexHookup(t *testing.T) {
	log.Printf("Hookup Succeeded - testing lexing")
}

func Test_Next(t *testing.T) {
	input := "abcd"
	l := lex("test", input)

	c := l.next()
	if c != 'a' {
		t.Error("next() not as expected")
	}

	c = l.next()
	if c != 'b' {
		t.Error("next() not as expected")
	}

	c = l.next()
	if c != 'c' {
		t.Error("next() not as expected")
	}

	c = l.next()
	if c != 'd' {
		t.Error("next() not as expected")
	}

	c = l.next()
	if c != eof {
		t.Error("next() not eof as expected")
	}
}

func Test_Peek(t *testing.T) {
	input := "abcd"
	l := lex("test", input)

	c := l.next()
	if c != 'a' {
		t.Error("next() not as expected")
	}

	c = l.next()
	if c != 'b' {
		t.Error("next() not as expected")
	}

	c = l.peek()
	if c != 'c' {
		t.Error("peek() not as expected")
	}

	c = l.peek()
	if c != 'c' {
		t.Error("peek() not as expected")
	}

	l.backup()
	c = l.next()
	if c != 'b' {
		t.Error("backup() not as expected")
	}
}

func Test_Accept(t *testing.T) {
	input := "a1"
	l := lex("test_accept", input)

	accepted := l.accept("123")
	if accepted {
		t.Error("accept() not as expected")
	}

	accepted = l.accept("abc")
	if !accepted {
		t.Error("accept() not as expected")
	}

	accepted = l.accept("123")
	if !accepted {
		t.Error("accept() not as expected")
	}
}

func Test_AcceptRun(t *testing.T) {
	input := "aabbcc1234"
	l := lex("test accept run", input)
	l.acceptRun("abcd")
	l.emit(itemText)

	item := l.nextItem()

	if item.val != "aabbcc" {
		t.Error("accept run not as expected")
	}
}

func Test_Indent(t *testing.T) {
	input := " "
	l := lex("test indent", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, " ", t)

	// test with indent
	input = "\t\t"
	l = lex("test indent 2", input)
	lexeme = l.nextItem()
	testLexeme(lexeme, itemIndentation, "\t\t", t)
}

func Test_Doctype(t *testing.T) {
	input := "!!!\nTest"
	l := lex("test doctype", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemDoctype, "", t)

	input = "!!! Strict\nmore content"
	l = lex("test strict", input)
	lexeme = l.nextItem()
	testLexeme(lexeme, itemDoctype, "Strict", t)
}

func Test_SimpleTag(t *testing.T) {
	input := "%simple_tag"
	l := lex("simple tag", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "simple_tag", t)
}

func Test_TagWithClassAndId(t *testing.T) {
	input := " %aTag#Id.class"
	l := lex("tag with id and class", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, " ", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "aTag", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemId, "Id", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemClass, "class", t)
}

func Test_Attribute(t *testing.T) {
	input := "%tag1{\"class\":\"yes\", width: 5px} content\n%tag2"
	l := lex("tag with attribute", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "tag1", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemAttributeName, "class", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemAttributeValue, "yes", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemAttributeName, "width", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemAttributeValue, "5px", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemText, "content", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemNewline, "\n", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "tag2", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemEOF, "", t)
}

func Test_Escape(t *testing.T) {
	input := "\\#1 test for backslashes"
	l := lex("escape test", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemText, "#1 test for backslashes", t)
}

func Test_HtmlComments(t *testing.T) {
	input := "\t/ %test test comment"
	l := lex("escape test", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "\t", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemHtmlComment, "/", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "test", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemText, "test comment", t)

}

func Test_SingleTag(t *testing.T) {
	input := "%test"
	l := lex("single tag", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "test", t)
}

func Test_TwoTags(t *testing.T) {
	input := "%tag1\n%tag2"
	l := lex("two tags", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "tag1", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemNewline, "\n", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "tag2", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemEOF, "", t)
}

func Test_Equals(t *testing.T) {
	input := "= code_output"
	l := lex("equals", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemCodeOutput, "code_output", t)
}

func Test_TagEquals(t *testing.T) {
	input := "%p= array[i]"
	l := lex("tag_code test", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemTag, "p", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemCodeOutput, "array[i]", t)
}

func Test_Dash(t *testing.T) {
	input := "- fmt.Println()"
	l := lex("dash", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemCodeExecution, "fmt.Println()", t)
}

func Test_DataType(t *testing.T) {
	input := "@data_type: string"
	l := lex("data_type", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemDataType, "string", t)
}

func Test_SingleImport(t *testing.T) {
	input := "@import ( \"fmt\" )"
	l := lex("single_import", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemImport, "fmt", t)
}

func Test_MultipleImports(t *testing.T) {
	input := "@import (\n\t\"fmt\"\n\t\"strings\"\n)"
	l := lex("data_type", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemImport, "fmt", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemImport, "strings", t)
}

func Test_TextStartsWithHyperlink(t *testing.T) {
	input := "  <a href='/'>Hey</a>"
	l := lex("hyperlink", input)

	lexeme := l.nextItem()
	testLexeme(lexeme, itemIndentation, "  ", t)

	lexeme = l.nextItem()
	testLexeme(lexeme, itemText, "<a href='/'>Hey</a>", t)
}

func testLexeme(l lexeme, expectedType lexItemType, expectedVal string, t *testing.T) {
	if l.typ != expectedType {
		t.Errorf("lexeme item type (%q) not as expected (%q)", l.typ, expectedType)
	}

	if l.val != expectedVal {
		t.Errorf("lexeme val (%q) not as expected (%q)", l.val, expectedVal)
	}
}
