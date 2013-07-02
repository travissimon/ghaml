package main

import (
	"container/list"
	"log"
)

type ParseContext struct {
	dataType string
	imports  []string
}

type GhamlParser struct {
	name    string
	lineNo  int
	lexer   *lexer
	root    *Node
	context *ParseContext

	// used during construction of the tree
	tags *stack
}

// initialises and returns a new Ghaml parser
func NewParser(name, input string) *GhamlParser {
	g := &GhamlParser{
		name:   name,
		lexer:  lex(name, input),
		lineNo: 0,
		tags:   new(stack),
		context: &ParseContext{
			dataType: "interface{}",
			imports:  make([]string, 0),
		},
	}

	g.root = buildNode("root")
	tagIndent := &tagIndentation{
		node:        g.root,
		indentLevel: 0,
	}
	g.tags.push(tagIndent)

	return g
}

// stack to track tag indentation for closing tags
type tagIndentation struct {
	node        *Node
	indentLevel int
}

// Represents ... wait for it... name and value strings
type nameValueStr struct {
	name  string
	value string
}

// A parse node (roughly translates directly to an html tag)
type Node struct {
	name        string
	text        string
	id          *nameValueStr // treat id and class differently so that we can 
	class       *nameValueStr // write them to the front of the attr list
	attributes  *list.List
	children    *list.List
	selfClosing bool
}

// appends a node to the children list
func (n *Node) appendNode(NodeToAdd *Node) {
	n.children.PushBack(NodeToAdd)
}

// sets the id of the node
func (n *Node) setId(id string) {
	n.id = &nameValueStr{
		name:  "id",
		value: id,
	}
}

// adds a class to the current node
func (n *Node) addClass(class string) {
	if n.class == nil {
		n.class = &nameValueStr{
			name:  "class",
			value: "",
		}
	}

	if len(n.class.value) > 0 {
		n.class.value += ", "
	}

	n.class.value += class
}

// adds an attribute to this node
func (n *Node) addAttribute(attr *nameValueStr) {
	n.attributes.PushBack(attr)
}

// adds an import to the context
func (g *GhamlParser) addImport(imp string) {
	g.context.imports = append(g.context.imports, imp)
}

// sets the 'data' parameter's type
func (g *GhamlParser) setDataType(dataType string) {
	g.context.dataType = dataType
}

// gets the current indent tag
func (g *GhamlParser) getCurrentTagIndent() *tagIndentation {
	return g.tags.peek().(*tagIndentation)
}

// gets the current parse node
func (g *GhamlParser) getCurrentNode() *Node {
	return g.getCurrentTagIndent().node
}

// gets the number of spaces in the current indent
func (g *GhamlParser) getCurrentIndentation() int {
	return g.getCurrentTagIndent().indentLevel
}

// initiates the parsing process
func (g *GhamlParser) Parse() {
Loop:
	for {
		lexeme := g.lexer.nextItem()

		switch lexeme.typ {
		case itemAttributeName:
			g.handleAttribute(lexeme.val)
		case itemImport:
			g.addImport(lexeme.val)
		case itemDataType:
			g.setDataType(lexeme.val)
		case itemDoctype:
			g.handleDoctype(lexeme)
		case itemCodeOutput:
			g.handleCodeOutput(lexeme)
		case itemCodeExecution:
			g.handleCodeExecution(lexeme)
		case itemIndentation:
			indentation := lexeme.val
			nextLexeme := g.lexer.nextItem()
			g.parseLineStart(indentation, nextLexeme)
		case itemText:
			g.getCurrentNode().text += lexeme.val
		case itemId:
			g.getCurrentNode().setId(lexeme.val)
		case itemClass:
			g.getCurrentNode().addClass(lexeme.val)
		case itemEOF:
			break Loop
		}
	}
}

// parses an attribute
func (g *GhamlParser) handleAttribute(attributeName string) {
	nextLexeme := g.lexer.nextItem()
	if nextLexeme.typ != itemAttributeValue {
		doError(g.lexer.lineNumber(), "expected attribute value, received "+nextLexeme.val)
	}
	value := nextLexeme.val
	attr := &nameValueStr{
		name:  attributeName,
		value: value,
	}
	g.getCurrentNode().addAttribute(attr)
}

// parses a doctype (!!!)
func (g *GhamlParser) handleDoctype(l lexeme) {
	n := buildNode("doctype")
	n.text = l.val
	n.selfClosing = true
	g.getCurrentNode().appendNode(n)
}

// parses a code output token (= ...)
func (g *GhamlParser) handleCodeOutput(l lexeme) {
	g.buildCodeNode("code_output", l)
}

// parses a code execution token (- ...)
func (g *GhamlParser) handleCodeExecution(l lexeme) {
	g.buildCodeNode("code_execution", l)
}

func (g *GhamlParser) buildCodeNode(nodeName string, l lexeme) {
	n := buildNode(nodeName)
	n.text = l.val
	g.getCurrentNode().appendNode(n)
}

// parses significant whitespace, and then handles the first node on the line
func (g *GhamlParser) parseLineStart(indentation string, firstItem lexeme) {
	indentationAmnt := len(indentation)

	for g.tags.count() > 1 && g.getCurrentIndentation() >= indentationAmnt {
		g.tags.pop()
	}

	if firstItem.typ == itemText {
		g.getCurrentNode().text += " " + firstItem.val
		return
	}

	var firstNode *Node
	if firstItem.typ == itemTag {
		firstNode = buildNode(firstItem.val)
	} else if firstItem.typ == itemCodeOutput {
		firstNode = buildNode("code_output")
		firstNode.text = firstItem.val
	} else if firstItem.typ == itemCodeExecution {
		firstNode = buildNode("code_execution")
		firstNode.text = firstItem.val
	} else {
		firstNode = buildNode("div")
	}

	tagIndentation := buildTagIndentation(firstNode, indentation)
	g.getCurrentNode().appendNode(tagIndentation.node)
	g.tags.push(tagIndentation)

	// if this was an implicit div, we stll need to add the id or class data
	switch firstItem.typ {
	case itemId:
		g.getCurrentNode().setId(firstItem.val)
	case itemClass:
		g.getCurrentNode().addClass(firstItem.val)
	}
}

func doError(lineNo int, msg string) {
	log.Printf("error line %q: %q", lineNo, msg)
}

func buildTagIndentation(n *Node, indentation string) *tagIndentation {
	return &tagIndentation{
		node:        n,
		indentLevel: len(indentation),
	}
}

func buildNode(name string) *Node {
	return &Node{
		name:        name,
		text:        "",
		id:          nil,
		class:       nil,
		attributes:  new(list.List),
		children:    new(list.List),
		selfClosing: false,
	}
}

// utility function to create a string dump of the content
func (g *GhamlParser) dumpNodes() string {
	result := ""

	for n := g.root.children.Front(); n != nil; n = n.Next() {
		result = dumpNode(n.Value.(*Node), result, 0)
	}

	return result
}

// recursive function to dump a node's content
func dumpNode(nd *Node, str string, indent int) string {
	if len(str) > 0 {
		str += "\n"
	}

	for i := 0; i < indent; i++ {
		str += "\t"
	}

	str += nd.name

	if nd.id != nil {
		str += " " + getAttrStr(nd.id)
	}

	if nd.class != nil {
		str += " " + getAttrStr(nd.class)
	}

	for attrEl := nd.attributes.Front(); attrEl != nil; attrEl = attrEl.Next() {
		attr := attrEl.Value.(*nameValueStr)
		str += " " + getAttrStr(attr)
	}

	if len(nd.text) > 0 {
		str += " (" + nd.text + ")"
	}

	for n := nd.children.Front(); n != nil; n = n.Next() {
		str = dumpNode(n.Value.(*Node), str, indent+1)
	}

	return str
}

// returns a string representation of an attribute
func getAttrStr(nv *nameValueStr) string {
	return nv.name + "='" + nv.value + "'"
}
