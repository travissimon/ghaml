package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type badMarshaler struct{}

func (x *badMarshaler) MarshalJSON() ([]byte, error) {
	// Keys in valid JSON must be double quoted as must all strings.
	return []byte("{ foo: 'not quite valid JSON' }"), nil
}

type goodMarshaler struct{}

func (x *goodMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`{ "<foo>": "O'Reilly" }`), nil
}

type TestDataType struct {
	F, T                     bool
	C, G, H                  string
	A, E                     []string
	B, M                     json.Marshaler
	N                        int
	Z                        *int
	constant                 string
	urlStartRel              string
	urlStartAbsOk            string
	protocolRelativeURLStart string
}

func getTestData() *TestDataType {
	return &TestDataType{
		F:                        false,
		T:                        true,
		C:                        "<Cincinatti>",
		G:                        "<Goodbye>",
		H:                        "<Hello>",
		A:                        []string{"<a>", "<b>"},
		E:                        []string{},
		N:                        42,
		B:                        &badMarshaler{},
		M:                        &goodMarshaler{},
		Z:                        nil,
		constant:                 "a<b",
		urlStartRel:              "/foo/bar?a=b&c=d",
		urlStartAbsOk:            "http://example.com/foo/bar?a=b&c=d",
		protocolRelativeURLStart: "//example.com:8000/foo/bar?a=b&c=d",
		pathRelativeURLStart:     "/javascript:80/foo/bar",
		dangerousURLStart:        "javascript:alert(%22pwned%22)",
		nonHierURL:               "mailto:Muhammed \"The Greatest\" Ali <m.ali@example.com>",
	}
}

func handleEscapeRequest(w http.ResponseWriter, r *http.Request) {
	testData := getTestData()
	wr := NewEscapingWriter()
	wr.SetData(testData)
	wr.Execute(w, r)
}

func main() {
	url := "localhost:8000"
	fmt.Printf("Listening on: http://%s/\n", url)

	http.HandleFunc("/", handleEscapeRequest)
	http.ListenAndServe(url, nil)
}
