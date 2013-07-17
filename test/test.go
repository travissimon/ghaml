package test

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"net/http"
	"text/template"
)

func NewTestWriter(data string) (*TestWriter) {
	wr := &TestWriter {
		data: data,
	}
	
	return wr
}

type TestWriter struct {
	data string
}

var TestHtml = [...]string{
`<html>
	<head>
		<title>
			`,
			`
		</title>
	</head>
	<body>
		<h1>
			`,
			`
			<div></div>
		</h1>
		<div>
			 This is child content for the div above. Note that HAML is space-sensitive, so all text indented at this
			level is encased in the div.
		</div>
		<div id="id_div"> You can use the # operator as a shortcut to create a div with the given id.</div>
		<div class="implicit_class">
			 The .operator (think of the '.' css selector') lets you create a div with the given class. For example
			this text will be wrapped in a div that looks like
		</div>
		`,
		`
		<div></div>
		<ul type="disc">
			`,
			`
			<li>
				`,
				`
			</li>
			`,
			`
		</ul>
	</body>
</html>
`,
}

func (wr TestWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *TestWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data string) {
	fmt.Fprint(w, TestHtml[0])
	fmt.Fprint(w, template.HTMLEscaper("Hello, ", data))
	fmt.Fprint(w, TestHtml[0])
	fmt.Fprint(w, template.HTMLEscaper("Hello, ", data))
	fmt.Fprint(w, TestHtml[0])
	fmt.Fprint(w, "Unescaped (and dangerous) output: <i>", data, "</i>")
	fmt.Fprint(w, TestHtml[0])
	for i := 0; i < 10; i++ {
		fmt.Fprint(w, TestHtml[1])
		fmt.Fprint(w, "Item: ", i)
		fmt.Fprint(w, TestHtml[0])
	}
}
