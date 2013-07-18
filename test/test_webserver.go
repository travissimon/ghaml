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
	W                        string
	Z                        *int
	Constant                 string
	UrlStartRel              string
	UrlStartAbsOk            string
	ProtocolRelativeURLStart string
	PathRelativeURLStart     string
	DangerousURLStart        string
	NonHierURL               string
	UrlPath                  string
	JsRe                     string
	Blank                    string
	Ltr                      string
	Left                     string
	Right                    string
	Expression               string
	Selector                 string
	ObfuscatedExpression     string
	ObfuscatedMozBinding     string
	Img                      string
	StyleURLBadProtocol      string
	Event                    string
	Code                     string
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
		W:                        `&iexcl;<b class="foo">Hello</b>, <textarea>O'World</textarea>!`,
		Z:                        nil,
		Constant:                 "a<b",
		UrlStartRel:              "/foo/bar?a=b&c=d",
		UrlStartAbsOk:            "http://example.com/foo/bar?a=b&c=d",
		ProtocolRelativeURLStart: "//example.com:8000/foo/bar?a=b&c=d",
		PathRelativeURLStart:     "/javascript:80/foo/bar",
		DangerousURLStart:        "javascript:alert(%22pwned%22)",
		NonHierURL:               "mailto:Muhammed \"The Greatest\" Ali <m.ali@example.com>",
		UrlPath:                  "javascript:80",
		JsRe:                     "foo+bar",
		Blank:                    "",
		Ltr:                      "ltr",
		Left:                     "left",
		Right:                    "right",
		Expression:               "expression(alert(1337))",
		Selector:                 "p",
		ObfuscatedExpression:     "  e\\78preS\x00Sio/**/n(alert(1337))",
		ObfuscatedMozBinding:     "  -mo\\7a-B\x00I/**/nding(alert(1337))",
		Img:                      "O'Reilly Animal(1)<2>.png",
		StyleURLBadProtocol:      "javascript:alert(1337)",
		Event:                    "onchange",
		Code:                     "doEvil()",
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
