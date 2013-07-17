package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Parsing method (heavily) inspired by Go's template parser.
// Thanks be to Rob Pike. see:
// http://blog.golang.org/2011/09/two-go-talks-lexical-scanning-in-go-and.html
// http://golang.org/src/pkg/text/template/parse/lex.go

// lexeme is 
type lexeme struct {
	typ lexItemType
	val string
}

// String() for item
func (l lexeme) String() string {
	if len(l.val) > 10 {
		return fmt.Sprintf("{%s, %.10q...}", l.typ, l.val)
	}
	return fmt.Sprintf("{%s, %q}", l.typ, l.val)
}

// types of lex items
type lexItemType int

const (
	itemError lexItemType = iota
	itemNewline
	itemIndentation
	itemDoctype
	itemText
	itemTag
	itemId
	itemClass
	itemAttributeName
	itemAttributeValue
	itemHtmlComment
	itemEOF
	itemCodeOutput
	itemRawCodeOutput
	itemCodeExecution
	itemDataType
	itemImport
	itemPackage
)

// pretty print items
var itemName = map[lexItemType]string{
	itemError:          "error",
	itemNewline:        "newline",
	itemIndentation:    "indentation",
	itemDoctype:        "doctype",
	itemText:           "text",
	itemTag:            "tag",
	itemId:             "id",
	itemClass:          "class",
	itemAttributeName:  "attribute name",
	itemAttributeValue: "attribute value",
	itemHtmlComment:    "Html comment",
	itemEOF:            "EOF",
	itemCodeOutput:     "code output",
	itemRawCodeOutput:  "raw code output",
	itemCodeExecution:  "code execution",
	itemDataType:       "data type",
	itemImport:         "import",
	itemPackage:        "package",
}

func (item lexItemType) String() string {
	s := itemName[item]
	if s == "" {
		return fmt.Sprintf("item%d", int(item))
	}
	return s
}

// state function represents 'current state' combined with 'next action'
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner
type lexer struct {
	name    string      // name of the input (for error reporting)
	input   string      // the string being scanned
	state   stateFn     // next lexing function
	pos     int         // current position in the input string
	start   int         // start position of this item
	width   int         // length of the last input rune
	lexemes chan lexeme // channel of scanned lexemes
}

// create a new lexer
func lex(name, input string) *lexer {
	l := &lexer{
		name:    name,
		input:   input,
		state:   lexLineStart,
		lexemes: make(chan lexeme, 2), // Two items of buffering is sufficient for all state functions
	}
	return l
}

// represent EOF when we're parsing an input string
const eof = -1

// next returns the next rune in the string
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// backup steps back one rune. Should only be called once per next()
func (l *lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune
func (l *lexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return r
}

// emit an item back to the client
func (l *lexer) emit(t lexItemType) {
	l.lexemes <- lexeme{t, l.previewCurrent()}
	l.start = l.pos
}

func (l *lexer) previewCurrent() string {
	return l.input[l.start:l.pos]
}

// skips the pending input
func (l *lexer) ignore() {
	l.start = l.pos
}

// accepts a rune if it's in the validRunes string
func (l *lexer) accept(validRunes string) bool {
	if strings.IndexRune(validRunes, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// accepts a series of runes that match the validRunes characters
func (l *lexer) acceptRun(validRunes string) {
	for strings.IndexRune(validRunes, l.next()) >= 0 {
	}
	l.backup()
}

// accepts a series of runes that are not in the invalidRunes characters
func (l *lexer) acceptRunUntil(invalidRunes string) {
	for {
		rune := l.next()
		if rune == eof {
			break
		}

		isFound := strings.IndexRune(invalidRunes, rune)
		if isFound >= 0 {
			break
		}
	}
	l.backup()
}

// skips all spaces and tabs from the current position
func (l *lexer) skipSpacesAndTabs() {
	l.acceptRun(" \t")
	l.ignore()
}

// which line are we currently on?
func (l *lexer) lineNumber() int {
	return 1 + strings.Count(l.input[:l.pos], "\n")
}

// error returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.lexemes <- lexeme{itemError, fmt.Sprintf(format, args...)}
	return nil
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
}

// nextItem returns the next item from the input
func (l *lexer) nextItem() lexeme {
	for {
		select {
		case lexeme := <-l.lexemes:
			if lexeme.typ == itemError {
				panic("ERROR: " + lexeme.val + "\n")
			}
			return lexeme
		default:
			l.state = l.state(l)
		}
	}
	panic("Not reached")
}

func lexLineStart(l *lexer) stateFn {
	switch l.peek() {
	case eof:
		l.emit(itemEOF)
		return nil
	case '!':
		return lexDoctype
	case '@':
		return lexMetadata
	case '\r', '\n':
		l.acceptRun("\n\r")
		l.ignore()
		return lexLineStart
	}

	return lexIndent
}

// lexes the end of line
func lexLineEnd(l *lexer) stateFn {
	l.skipSpacesAndTabs()

	switch l.peek() {
	case eof:
		l.emit(itemEOF)
		return nil
	case '\n', '\r':
		return lexNewline
	}

	l.errorf("Expected end of line, received %q", l.peek())
	return nil
}

// lexes the start of a line
func lexIndent(l *lexer) stateFn {
	l.acceptRun(" \t")
	l.emit(itemIndentation)
	return lexContentStart
}

// lexes doc type declarations
func lexDoctype(l *lexer) stateFn {
	// skip "!!!" and trailing space
	l.acceptRun("! ")
	l.ignore()

Loop:
	for {
		switch l.next() {
		case eof, '\n':
			break Loop
		}
	}
	l.backup()
	l.emit(itemDoctype)
	return lexLineEnd
}

// lexes the start of content (after significant whitespace)
func lexContentStart(l *lexer) stateFn {
	switch l.peek() {
	case eof, '\n':
		return lexLineEnd
	case '\\':
		l.next()
		l.ignore()
		return lexText
	case '%':
		return lexTag
	case '.':
		return lexClass
	case '#':
		return lexId
	case '/':
		return lexHtmlComment
	case '=':
		return lexCodeOutput
	case '|':
		return lexRawCodeOutput
	case '-':
		return lexCodeExecution
	}
	return lexText
}

// lexes the 'name' of something
func (l *lexer) lexIdentifier(lexType lexItemType) stateFn {
	specialChars := " #.%${}:'\"/?=|-"
	l.accept(specialChars)
	l.ignore()
	l.acceptRunUntil(specialChars + "\n\r")
	l.emit(lexType)

	switch l.peek() {
	case eof, '\r', '\n':
		return lexLineEnd
	case ' ', '\t':
		return lexText
	case '#':
		return lexId
	case '.':
		return lexClass
	case '{':
		return lexAttribute
	case '=':
		return lexCodeOutput
	case '|':
		return lexRawCodeOutput
	case '-':
		return lexCodeExecution
	default:
	}
	return lexLineEnd
}

// lexes a tag declaration (e.g. %tag)
func lexTag(l *lexer) stateFn {
	return l.lexIdentifier(itemTag)
}

// lexes an id declaration (e.g. #id)
func lexId(l *lexer) stateFn {
	return l.lexIdentifier(itemId)
}

// lexes a class declaration (e.g. .class)
func lexClass(l *lexer) stateFn {
	return l.lexIdentifier(itemClass)
}

// lexes free-form text
func lexText(l *lexer) stateFn {
	l.skipSpacesAndTabs()
Loop:
	for {
		switch l.next() {
		case eof, '\n', '\r':
			break Loop
		}
	}
	l.backup()
	l.emit(itemText)

	return lexLineEnd
}

// lexes a tag attribute
func lexAttribute(l *lexer) stateFn {
	l.accept(",{")
	l.skipSpacesAndTabs()
	l.accept("\"")
	l.ignore()
	l.acceptRunUntil("\":")
	l.emit(itemAttributeName)
	l.accept("\"")
	l.skipSpacesAndTabs()
	if !l.accept(":") {
		return l.errorf("Unexpected rune %q. Expecting ':'")
	}
	l.skipSpacesAndTabs()
	l.accept("\"")
	l.ignore()
	l.acceptRunUntil("\"},")
	l.emit(itemAttributeValue)
	l.accept("\"")
	l.ignore()

	c := l.next()
	l.ignore()
	switch c {
	case ',':
		return lexAttribute
	case '}':
		return lexText
	}
	return l.errorf("Expected end of attribute, received %q", l.peek())
}

// lexes a newline
func lexNewline(l *lexer) stateFn {
	l.acceptRun("\r\n")
	l.emit(itemNewline)
	return lexLineStart
}

// lexes markup for an html comment
func lexHtmlComment(l *lexer) stateFn {
	l.accept("/")
	l.emit(itemHtmlComment)
	l.acceptRun(" \t")
	l.ignore()
	return lexContentStart
}

// lexes markup for output to be escaped and printed
func lexCodeOutput(l *lexer) stateFn {
	l.accept("=")
	l.skipSpacesAndTabs()
	l.ignore()
Loop:
	for {
		switch l.next() {
		case eof, '\n', '\r':
			break Loop
		}
	}
	l.backup()
	l.emit(itemCodeOutput)
	return lexLineEnd
}

// lexes markup for output to be printed without XSS safety
func lexRawCodeOutput(l *lexer) stateFn {
	l.accept("|")
	l.skipSpacesAndTabs()
	l.ignore()
Loop:
	for {
		switch l.next() {
		case eof, '\n', '\r':
			break Loop
		}
	}
	l.backup()
	l.emit(itemRawCodeOutput)
	return lexLineEnd
}

// lexes markup for go code to be executed but not printed
func lexCodeExecution(l *lexer) stateFn {
	l.accept("-")
	l.skipSpacesAndTabs()
	l.ignore()
Loop:
	for {
		switch l.next() {
		case eof, '\n', '\r':
			break Loop
		}
	}

	l.backup()
	l.emit(itemCodeExecution)
	return lexLineEnd
}

// lexes ghaml-specific meta-data
func lexMetadata(l *lexer) stateFn {
	l.acceptRunUntil(": (\n\r")

	metadataType := l.previewCurrent()
	switch metadataType {
	case "@data_type":
		l.ignore()
		return lexDatatype
	case "@import":
		l.ignore()
		return lexImports
	case "@package":
		l.ignore()
		return lexPackage
	}

	l.errorf("Expected @data_type or @import. Received %q", metadataType)
	return nil
}

// lexes the ghaml-specific datatype specifier
func lexDatatype(l *lexer) stateFn {
	l.acceptRun(": ")
	l.ignore()
	l.acceptRunUntil(" \n\r")
	l.emit(itemDataType)
	return lexLineEnd
}

// lexes the ghaml-specific 'import' statements
func lexImports(l *lexer) stateFn {
	l.acceptRun(" (")
	l.ignore()
	for {
		l.acceptRun(" \n\r\t")
		switch l.peek() {
		case ')':
			l.accept(")")
			l.skipSpacesAndTabs()
			l.ignore()
			return lexLineEnd
		case '"':
			l.accept("\"")
			l.ignore()
			l.acceptRunUntil("\"")
			l.emit(itemImport)
			l.accept("\"")
		default:
			l.errorf("Failed to parse import. Expected '(' or '\"', found: %q", l.peek())
			break
		}
	}
	return nil
}

func lexPackage(l *lexer) stateFn {
	l.acceptRun(": ")
	l.ignore()
	l.acceptRunUntil(" \n\r")
	l.emit(itemPackage)
	return lexLineEnd
}
