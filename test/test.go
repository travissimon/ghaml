package main

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"net/http"
)

func NewTestWriter() (*TestWriter) {
	wr := &TestWriter{}
	
	for idx, pattern := range TestTemplatePatterns {
		tmpl, err := template.New("TestTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		TestTemplates = append(TestTemplates, tmpl)
	}
	return wr
}

type TestWriter struct {
	data string
}

func (wr *TestWriter) SetData(data string) {
	wr.data = data
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
		 | "Unescaped (and dangerous) output: <i>", data, "</i>"
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

var TestTemplatePatterns = []string{
	`Hello, {{.}}`,
	`Hello, {{.}}`,
}

var TestTemplates = make([]*template.Template, 0, len(TestTemplatePatterns))

func (wr TestWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *TestWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data string) {
	var err error = nil
	fmt.Fprint(w, TestHtml[0])
	err = TestTemplates[0].Execute(w, data)
	handleTestError(err)
	fmt.Fprint(w, TestHtml[1])
	err = TestTemplates[1].Execute(w, data)
	handleTestError(err)
	fmt.Fprint(w, TestHtml[2])
	for i := 0; i < 10; i++ {
		fmt.Fprint(w, TestHtml[3])
		fmt.Fprint(w, "Item: ", i)
		fmt.Fprint(w, TestHtml[4])
	}
	fmt.Fprint(w, TestHtml[5])
}

func handleTestError(err error) {
	if err != nil {fmt.Println(err)}}