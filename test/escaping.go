package main

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"net/http"
)

func NewEscapingWriter() (*EscapingWriter) {
	wr := &EscapingWriter{}
	
	for idx, pattern := range EscapingTemplatePatterns {
		tmpl, err := template.New("EscapingTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		EscapingTemplates = append(EscapingTemplates, tmpl)
	}
	return wr
}

type EscapingWriter struct {
	data *TestDataType
}

func (wr *EscapingWriter) SetData(data *TestDataType) {
	wr.data = data
}

var EscapingHtml = [...]string{
`<html>
	<body>
		<div id="if">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
				`,
				`
				`,
				`
			</div>
		</div>
		<div id="else">
			<div class="exepected">
				`,
				`
			</div>
			<div></div>
			<div class="actual">
				`,
				`
				`,
				`
				`,
				`
			</div>
		</div>
	</body>
</html>
`,
}

var EscapingTemplatePatterns = []string{
	`Hello,`,
	`{{.C}}!`,
	`{{.H}}`,
	`{{.G}}`,
}

var EscapingTemplates = make([]*template.Template, 0, len(EscapingTemplatePatterns))

func (wr EscapingWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *EscapingWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data *TestDataType) {
	fmt.Fprint(w, EscapingHtml[0])
	fmt.Fprintf(w, "Hello, &lt;Cincinatti&gt;!")
	fmt.Fprint(w, EscapingHtml[1])
	if data.T {
		fmt.Fprint(w, EscapingHtml[2])
		EscapingTemplates[0].Execute(w, data)
	}
	fmt.Fprint(w, EscapingHtml[3])
	EscapingTemplates[1].Execute(w, data)
	fmt.Fprint(w, EscapingHtml[4])
	fmt.Fprintf(w, "&lt;Goodbye&gt;")
	fmt.Fprint(w, EscapingHtml[5])
	if data.F {
		fmt.Fprint(w, EscapingHtml[6])
		EscapingTemplates[2].Execute(w, data)
	} else {
		fmt.Fprint(w, EscapingHtml[7])
		EscapingTemplates[3].Execute(w, data)
	}
	fmt.Fprint(w, EscapingHtml[8])
}
