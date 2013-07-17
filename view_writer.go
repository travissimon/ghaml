package main

import (
	"bytes"
	"github.com/travissimon/formatting"
	"io"
	"strings"
)

type ViewWriter struct {
	context         *ParseContext
	rootNode        *Node
	writer          io.Writer
	destinationName string
}

func NewViewWriter(wr io.Writer, context *ParseContext, rootNode *Node, writerName string) *ViewWriter {
	vw := &ViewWriter{
		context:         context,
		rootNode:        rootNode,
		writer:          wr,
		destinationName: writerName,
	}

	return vw
}

func (w *ViewWriter) WriteView() {
	src := formatting.NewIndentingWriter(w.writer)

	htmlArr, srcOut, escapedHtml := w.processNodes()

	src.Printf("package %s\n", w.context.pkg)
	src.Println("")
	src.Println("// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN")
	src.Println("// EDIT THE .haml FILE INSTEAD")
	src.Println("")
	src.Println("import (")
	src.IncrIndent()
	src.Println("\"fmt\"")
	src.Println("\"net/http\"")
	if escapedHtml {
		src.Println("\"text/template\"")
	}
	for _, imp := range w.context.imports {
		src.Printf("%q\n", imp)
	}
	src.DecrIndent()
	src.Println(")")
	src.Println("")
	src.Printf("func New%sWriter(data %s) (*%sWriter) {\n", w.destinationName, w.context.dataType, w.destinationName)
	src.IncrIndent()
	src.Printf("wr := &%sWriter {\n", w.destinationName)
	src.IncrIndent()
	src.Printf("data: data,\n")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Println("return wr")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("type %sWriter struct {\n", w.destinationName)
	src.IncrIndent()
	src.Printf("data %s\n", w.context.dataType)
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("var %sHtml = [...]string{\n", w.destinationName)
	src.Println(htmlArr)
	src.Println("}")
	src.Println("")
	src.Printf("func (wr %sWriter) Execute(w http.ResponseWriter, r *http.Request) {\n", w.destinationName)
	src.IncrIndent()
	src.Println("wr.ExecuteData(w, r, wr.data)")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("func (wr *%sWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data %s) {\n", w.destinationName, w.context.dataType)

	// output from processNodes
	// This is calls to htmlArray and code generated Prints
	src.Print(srcOut)

	src.Println("}")
}

// processNodes generates code from the parsed haml nodes
// htmlArray is a string array of the raw html parts.
// src is the go source that calls html array, and then user defined code.
// usually looks like this:
//	fmt.Fprint(w, HtmlArray[0])
//	fmt.Fprint(w, "Custom string: ", data)
//  fmt.Fprint(w, HtmlArray[1]) ...
func (w *ViewWriter) processNodes() (htmlArray string, src string, escapedHtml bool) {
	htmlBuffer := bytes.NewBuffer(make([]byte, 0))
	htmlWriter := formatting.NewIndentingWriter(htmlBuffer)
	srcBuffer := bytes.NewBuffer(make([]byte, 0))
	srcWriter := formatting.NewIndentingWriter(srcBuffer)

	// initialise opening quote for htlmArray
	htmlWriter.Print("`")

	// initialise starting indent for src
	srcWriter.IncrIndent()

	currentHtmlIndex := 0
	esc := false
	for n := w.rootNode.children.Front(); n != nil; n = n.Next() {
		currentHtmlIndex, esc = w.writeNode(n.Value.(*Node), htmlWriter, srcWriter, currentHtmlIndex)
		escapedHtml = escapedHtml || esc
	}

	// close quote for html Array
	htmlWriter.Print("`,")

	htmlArray = htmlBuffer.String()
	src = srcBuffer.String()
	return
}

// Recursive function to write parsed HAML Nodes
// We have to return a bool indicating if we have escaped any HTML (XSS protection)
// so that we know if we need to include the templating library for that function
func (w *ViewWriter) writeNode(nd *Node, haml *formatting.IndentingWriter, src *formatting.IndentingWriter, currentHtmlIndex int) (htmlIndex int, escapedHtml bool) {
	escapedHtml = false
	if nd.name == "code_output" {
		escapedHtml = true
		htmlIndex = w.writeCodeOutput(nd, haml, src, currentHtmlIndex, true)
		return
	} else if nd.name == "raw_code_output" {
		htmlIndex = w.writeCodeOutput(nd, haml, src, currentHtmlIndex, false)
		return
	} else if nd.name == "code_execution" {
		htmlIndex = w.writeCodeExecution(nd, haml, src, currentHtmlIndex)
		return
	}

	haml.Printf("<%s", nd.name)
	if nd.id != nil {
		w.writeAttribute(nd.id, haml)
	}
	if nd.class != nil {
		w.writeAttribute(nd.class, haml)
	}

	for attrEl := nd.attributes.Front(); attrEl != nil; attrEl = attrEl.Next() {
		attr := attrEl.Value.(*nameValueStr)
		w.writeAttribute(attr, haml)
	}

	if nd.selfClosing {
		haml.Println(" />`)")
		return
	} else {
		haml.Print(">")
	}

	// Outputting text.

	// If tag only contains short text, add it on same line
	if w.canChildContentFitOnOneLine(nd) {
		haml.Printf("%s</%s>\n", nd.text, nd.name)
		return
	}

	// We either have long text, child tags or both
	// so we add it as indented child content
	haml.Println("")
	haml.IncrIndent()

	if len(nd.text) > 0 {
		w.writeLongText(nd.text, haml)
	}

	esc := false
	for n := nd.children.Front(); n != nil; n = n.Next() {
		currentHtmlIndex, esc = w.writeNode(n.Value.(*Node), haml, src, currentHtmlIndex)
		escapedHtml = escapedHtml || esc
	}

	haml.DecrIndent()
	haml.Printf("</%s>\n", nd.name)

	return
}

var TEXT_BREAK_LENGTH = 100

func (vw *ViewWriter) writeLongText(text string, w *formatting.IndentingWriter) {
	// create index of spaces in string
	spaces := vw.getWhitespaceIndicies(text)

	// split string on space less than MAX_STRING_LENGTH
	start := 0

	for _, idx := range spaces {
		distance := idx - start
		if distance > TEXT_BREAK_LENGTH {
			w.Println(text[start:idx])
			start = idx + 1
		}
	}
	w.Println(text[start:])
}

func (vw *ViewWriter) getWhitespaceIndicies(text string) []int {
	spaces := make([]int, 0, 255)

	for i, c := range text {
		switch c {
		case ' ', '\t', '\n':
			spaces = append(spaces, i)
		}
	}

	return spaces
}

// Child content can only fit on one line when there is short
// text and no child nodes
func (w *ViewWriter) canChildContentFitOnOneLine(nd *Node) bool {
	return len(nd.text) < TEXT_BREAK_LENGTH && nd.children.Len() == 0
}

func (w *ViewWriter) writeAttribute(attribute *nameValueStr, haml *formatting.IndentingWriter) {
	haml.Printf(" %s=\"%s\"", attribute.name, attribute.value)
}

func (w *ViewWriter) writeCodeOutput(nd *Node, haml *formatting.IndentingWriter, src *formatting.IndentingWriter, currentHtmlIndex int, escapeHtml bool) int {
	// end most recent haml output
	haml.Println("`,")
	// start next haml output (which will follow this code output
	haml.Println("`")

	// All call to write html from array 
	src.Printf("fmt.Fprint(w, %sHtml[%d])\n", w.destinationName, currentHtmlIndex)
	currentHtmlIndex++

	// add call to print output
	if escapeHtml {
		src.Printf("fmt.Fprint(w, template.HTMLEscaper(%s))\n", nd.text)
	} else {
		src.Printf("fmt.Fprint(w, %s)\n", nd.text)
	}

	for n := nd.children.Front(); n != nil; n = n.Next() {
		currentHtmlIndex, _ = w.writeNode(n.Value.(*Node), haml, src, currentHtmlIndex)
	}

	return currentHtmlIndex
}

func (w *ViewWriter) writeCodeExecution(nd *Node, haml *formatting.IndentingWriter, src *formatting.IndentingWriter, currentHtmlIndex int) int {
	// end most recent haml output
	haml.Println("`,")
	// start next haml output (which will follow this code output
	haml.Println("`")

	// All call to write html from array 
	src.Printf("fmt.Fprint(w, %sHtml[%d])\n", w.destinationName, currentHtmlIndex)
	currentHtmlIndex++

	// attempt to keep formatting across user code. 
	// Here we're checking to see if this is the end of a block statement
	// if so, we need to decrease indent
	first := getFirstChar(nd.text)
	if first == '}' {
		src.DecrIndent()
	}

	// add user's code
	src.Printf("%s\n", nd.text)

	// If user code ends in {, incr indent as they started a block statement
	last := getLastChar(nd.text)
	if last == '{' {
		src.IncrIndent()
	}

	for n := nd.children.Front(); n != nil; n = n.Next() {
		currentHtmlIndex, _ = w.writeNode(n.Value.(*Node), haml, src, currentHtmlIndex)
	}

	return currentHtmlIndex
}

func getFirstChar(s string) byte {
	trimmed := strings.TrimLeft(s, "\t ")
	return trimmed[0]
}

func getLastChar(s string) byte {
	trimmed := strings.TrimRight(s, "\t ")
	return trimmed[len(trimmed)-1]
}
