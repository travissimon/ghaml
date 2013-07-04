package main

import (
	"github.com/travissimon/formatting"
	"io"
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

	src.Println("package main")
	src.Println("")
	src.Println("// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN")
	src.Println("// EDIT THE .haml FILE INSTEAD")
	src.Println("")
	src.Println("import (")
	src.IncrIndent()
	src.Println("\"github.com/travissimon/formatting\"")
	src.Println("\"net/http\"")
	for i := range w.context.imports {
		src.Printf("%q\n", w.context.imports[i])
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
	src.Printf("func (wr %sWriter) Execute(w http.ResponseWriter, r *http.Request) {\n", w.destinationName)
	src.IncrIndent()
	src.Println("wr.ExecuteData(w, r, wr.data)")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("func (wr *%sWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data %s) {\n", w.destinationName, w.context.dataType)
	src.IncrIndent()
	src.Print("html := formatting.NewIndentingWriter(w)\n")
	src.Println("")
	src.Print("html.Print(\n`")

	// generate the parsed lines
	// The indenting for the parsed lines is separate to
	// the source formatting, so we create a new indenter
	hamlIndenter := formatting.NewIndentingWriter(src.Writer)
	for n := w.rootNode.children.Front(); n != nil; n = n.Next() {
		w.WriteNode(n.Value.(*Node), hamlIndenter, src)
	}

	src.Println("`)")
	src.DecrIndent()
	src.Println("}")
}

// Recursive function to write parsed HAML Nodes
func (w *ViewWriter) WriteNode(nd *Node, haml *formatting.IndentingWriter, src *formatting.IndentingWriter) {

	if nd.name == "code_output" {
		w.WriteCodeOutput(nd, haml, src)
		return
	} else if nd.name == "code_execution" {
		w.WriteCodeExecution(nd, haml, src)
		return
	}

	haml.Printf("<%s", nd.name)
	if nd.id != nil {
		w.WriteAttribute(nd.id, haml)
	}
	if nd.class != nil {
		w.WriteAttribute(nd.class, haml)
	}

	for attrEl := nd.attributes.Front(); attrEl != nil; attrEl = attrEl.Next() {
		attr := attrEl.Value.(*nameValueStr)
		w.WriteAttribute(attr, haml)
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

	for n := nd.children.Front(); n != nil; n = n.Next() {
		w.WriteNode(n.Value.(*Node), haml, src)
	}

	haml.DecrIndent()
	haml.Printf("</%s>\n", nd.name)
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

func (w *ViewWriter) WriteAttribute(attribute *nameValueStr, haml *formatting.IndentingWriter) {
	haml.Printf(" %s=\"%s\"", attribute.name, attribute.value)
}

func (w *ViewWriter) WriteCodeOutput(nd *Node, haml *formatting.IndentingWriter, src *formatting.IndentingWriter) {
	haml.Print("`")
	src.Println(")\n")
	src.Printf("html.Print(%s)\n", nd.text)
	haml.Println("")
	src.Print("html.Print(\n`\n")

	for n := nd.children.Front(); n != nil; n = n.Next() {
		w.WriteNode(n.Value.(*Node), haml, src)
	}
}

func (w *ViewWriter) WriteCodeExecution(nd *Node, haml *formatting.IndentingWriter, src *formatting.IndentingWriter) {
	haml.Print("`")
	src.Println(")\n")
	src.Printf("%s\n", nd.text)
	src.Print("html.Print(\n`")

	for n := nd.children.Front(); n != nil; n = n.Next() {
		w.WriteNode(n.Value.(*Node), haml, src)
	}
}
