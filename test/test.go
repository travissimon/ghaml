package main

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"github.com/travissimon/formatting"
	"net/http"
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

func (wr TestWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *TestWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data string) {
	html := formatting.NewIndentingWriter(w)

	html.Print(
`<html>
	<head>
		<title>
			`)

	html.Print(data)

	html.Print(
`
		</title>
	</head>
	<body>
		<h1>
			Test output
			<div></div>
		</h1>
		<div>This is a test. Hope it works out</div>
		<div></div>
		<div valign="top">supposed to be valign top here</div>
		<p>This is stuff:</p>
		<ul>
			`)

	for i := 0; i < 10; i++ {
	html.Print(
`<li>Travis is COOL</li>
			`)

	}
	html.Print(
`</ul>
	</body>
</html>
`)
}
