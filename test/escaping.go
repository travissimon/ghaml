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

func (wr *EscapingWriter) SetData(data interface{}) {
	wr.data = data.(*TestDataType)
}

var EscapingHtml = [...]string{
`<html>
	<body>
		<div></div>
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
		<div></div>
		<div id="rangeBody">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
				`,
				`
			</div>
		</div>
		<div id="constant">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="multipleAttrs">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlStartRel">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlStartAbsOk">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="protocolRelativeURLStart">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="pathRelativStart">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="dangerousURLStart">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="dangerousURLStart2">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="nonHierURL">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlPath">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlQuery">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="urlFragment">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsStrValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsNumericValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsBoolValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsNilValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsObjValue">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsObjValueScript">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsObjValueNotOverEscaped">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsStr">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="badMarshaler">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsMarshaler">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsStrNotUnderEscaped">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsRe">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="jsReBlank">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleBidiKeywordPassed">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleBidiPropNamePassed">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleExpressionBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleTagSelectorPassed">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleObfuscatedExpressionBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleObfuscatedMozBindingBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleURLQueryEncoded">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="styleURLBadProtocolBlocked">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInText">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInAttribute">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInScript">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="HtmlInRCDATA">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
				`,
				`
			</div>
		</div>
		<div id="DynamicAttributeName">
			<div class="expected">
				`,
				`
			</div>
			<div class="actual">
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
	`{{.}}`,
	`<a href="/search?q={{.Constant}}">`,
	`<a b=1 c={{.H}}>`,
	`<a href='{{.UrlStartRel}}'>`,
	`<a href='{{.UrlStartAbsOk}}'>`,
	`<a href='{{.ProtocolRelativeURLStart}}'>`,
	`<a href="{{.PathRelativeURLStart}}">`,
	`<a href='{{.DangerousURLStart}}'>`,
	`<a href=' {{.DangerousURLStart}}'>`,
	`<a href={{.NonHierURL}}>`,
	`<a href='http://{{.UrlPath}}/foo'>`,
	`<a href='/search?q={{.H}}'>`,
	`<a href='/faq#{{.H}}'>`,
	`<button onclick='alert({{.H}})'>`,
	`<button onclick='alert({{.N}})'>`,
	`<button onclick='alert({{.T}})'>`,
	`<button onclick='alert(typeof {{.Z}})'>`,
	`<button onclick='alert({{.A}})'>`,
	`<script>alert({{.A}})</script>"`,
	`<button onclick='alert({{.A}})'>`,
	`<button onclick='alert(&quot;{{.H}}&quot;)'>`,
	`<button onclick='alert(1/{{.B}}in numbers)'>`,
	`<button onclick='alert({{.M}})'>`,
	`<button onclick='alert({{.C}})'>`,
	`<button onclick='alert(/{{.JsRe}}/.test(""))'>`,
	`<script>alert(/{{.Blank}}/.test(""));</script>`,
	`<p style="dir: {{.Ltr}}">`,
	`<p style="border-{{.Left}}: 0; border-{{.Right}}: 1in">`,
	`<p style="width: {{"expression(alert(1337))"}}">`,
	`<style>{{.Selector}} { color: pink }</style>`,
	`<p style="width: {{.ObfuscatedExpression}}">`,
	`<p style="{{.ObfuscatedMozBinding}}: ...">`,
	`<p style="background: url(/img?name={{.Img}})">`,
	`<a style="background: url('{{.StyleURLBadProtocol}}')">`,
	`{{.W}}`,
	`<div title="{{.W}}">`,
	`<button onclick="alert({{.W}})">`,
	`<textarea>{{.W}}</textarea>`,
	`<input {{.Event}}="{{.Code}}">`,
}

var EscapingTemplates = make([]*template.Template, 0, len(EscapingTemplatePatterns))

func (wr EscapingWriter) Execute(w http.ResponseWriter, r *http.Request) {
	wr.ExecuteData(w, r, wr.data)
}

func (wr *EscapingWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data *TestDataType) {
	var err error = nil
	fmt.Fprint(w, EscapingHtml[0])
	fmt.Fprint(w, "Hello, &lt;Cincinatti&gt;!")
	fmt.Fprint(w, EscapingHtml[1])
	if data.T {
		fmt.Fprint(w, EscapingHtml[2])
		err = EscapingTemplates[0].Execute(w, data)
		handleEscapingError(err)
	}
	fmt.Fprint(w, EscapingHtml[3])
	err = EscapingTemplates[1].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[4])
	fmt.Fprint(w, "&lt;Goodbye&gt;")
	fmt.Fprint(w, EscapingHtml[5])
	if data.F {
		fmt.Fprint(w, EscapingHtml[6])
		err = EscapingTemplates[2].Execute(w, data)
		handleEscapingError(err)
	} else {
		fmt.Fprint(w, EscapingHtml[7])
		err = EscapingTemplates[3].Execute(w, data)
		handleEscapingError(err)
	}
	fmt.Fprint(w, EscapingHtml[8])
	fmt.Fprint(w, "&lt;a&gt;&lt;b&gt;")
	fmt.Fprint(w, EscapingHtml[9])
	for _, d := range data.A {
		fmt.Fprint(w, EscapingHtml[10])
		err = EscapingTemplates[4].Execute(w, d)
		handleEscapingError(err)
	}
	fmt.Fprint(w, EscapingHtml[11])
	fmt.Fprint(w, `<a href="/search?q=%27a%3cb%27">`)
	fmt.Fprint(w, EscapingHtml[12])
	err = EscapingTemplates[5].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[13])
	fmt.Fprint(w, "<a b=1 c=&lt;Hello&gt;>")
	fmt.Fprint(w, EscapingHtml[14])
	err = EscapingTemplates[6].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[15])
	fmt.Fprint(w, `<a href='/foo/bar?a=b&amp;c=d'>`)
	fmt.Fprint(w, EscapingHtml[16])
	err = EscapingTemplates[7].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[17])
	fmt.Fprint(w, `<a href='http://example.com/foo/bar?a=b&amp;c=d'>`)
	fmt.Fprint(w, EscapingHtml[18])
	err = EscapingTemplates[8].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[19])
	fmt.Fprint(w, `<a href='//example.com:8000/foo/bar?a=b&amp;c=d'>`)
	fmt.Fprint(w, EscapingHtml[20])
	err = EscapingTemplates[9].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[21])
	fmt.Fprint(w, `<a href="/javascript:80/foo/bar">`)
	fmt.Fprint(w, EscapingHtml[22])
	err = EscapingTemplates[10].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[23])
	fmt.Fprint(w, `<a href='#ZgotmplZ'>`)
	fmt.Fprint(w, EscapingHtml[24])
	err = EscapingTemplates[11].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[25])
	fmt.Fprint(w, `<a href=' #ZgotmplZ'>`)
	fmt.Fprint(w, EscapingHtml[26])
	err = EscapingTemplates[12].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[27])
	fmt.Fprint(w, `<a href=mailto:Muhammed%20%22The%20Greatest%22%20Ali%20%3cm.ali@example.com%3e>`)
	fmt.Fprint(w, EscapingHtml[28])
	err = EscapingTemplates[13].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[29])
	fmt.Fprint(w, `<a href='http://javascript:80/foo'>`)
	fmt.Fprint(w, EscapingHtml[30])
	err = EscapingTemplates[14].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[31])
	fmt.Fprint(w, `<a href='/search?q=%3cHello%3e'>`)
	fmt.Fprint(w, EscapingHtml[32])
	err = EscapingTemplates[15].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[33])
	fmt.Fprint(w, `<a href='/faq#%3cHello%3e'>`)
	fmt.Fprint(w, EscapingHtml[34])
	err = EscapingTemplates[16].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[35])
	fmt.Fprint(w, `<button onclick='alert(&#34;\u003cHello\u003e&#34;)'>`)
	fmt.Fprint(w, EscapingHtml[36])
	err = EscapingTemplates[17].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[37])
	fmt.Fprint(w, `<button onclick='alert( 42 )'>`)
	fmt.Fprint(w, EscapingHtml[38])
	err = EscapingTemplates[18].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[39])
	fmt.Fprint(w, `<button onclick='alert( true )'>`)
	fmt.Fprint(w, EscapingHtml[40])
	err = EscapingTemplates[19].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[41])
	fmt.Fprint(w, `<button onclick='alert(typeof null )'>`)
	fmt.Fprint(w, EscapingHtml[42])
	err = EscapingTemplates[20].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[43])
	fmt.Fprint(w, `<button onclick='alert([&#34;\u003ca\u003e&#34;,&#34;\u003cb\u003e&#34;])'>`)
	fmt.Fprint(w, EscapingHtml[44])
	err = EscapingTemplates[21].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[45])
	fmt.Fprint(w, `<script>alert(["\u003ca\u003e","\u003cb\u003e"])</script>`)
	fmt.Fprint(w, EscapingHtml[46])
	err = EscapingTemplates[22].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[47])
	fmt.Fprint(w, `<button onclick='alert([&#34;\u003ca\u003e&#34;,&#34;\u003cb\u003e&#34;])'>`)
	fmt.Fprint(w, EscapingHtml[48])
	err = EscapingTemplates[23].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[49])
	fmt.Fprint(w, `<button onclick='alert(&quot;\x3cHello\x3e&quot;)'>`)
	fmt.Fprint(w, EscapingHtml[50])
	err = EscapingTemplates[24].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[51])
	fmt.Fprint(w, `<button onclick='alert(1/ /* json: error calling MarshalJSON for type *template.badMarshaler: invalid character &#39;f&#39; looking for beginning of object key string */null in numbers)'>`)
	fmt.Fprint(w, EscapingHtml[52])
	err = EscapingTemplates[25].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[53])
	fmt.Fprint(w, `<button onclick='alert({&#34;\u003cfoo\u003e&#34;:&#34;O&#39;Reilly&#34;})'>`)
	fmt.Fprint(w, EscapingHtml[54])
	err = EscapingTemplates[26].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[55])
	fmt.Fprint(w, `<button onclick='alert(&#34;%3CCincinatti%3E&#34;)'>`)
	fmt.Fprint(w, EscapingHtml[56])
	err = EscapingTemplates[27].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[57])
	fmt.Fprint(w, `<button onclick='alert(/foo\x2bbar/.test(""))'>`)
	fmt.Fprint(w, EscapingHtml[58])
	err = EscapingTemplates[28].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[59])
	fmt.Fprint(w, `<script>alert(/(?:)/.test(""));</script>`)
	fmt.Fprint(w, EscapingHtml[60])
	err = EscapingTemplates[29].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[61])
	fmt.Fprint(w, `<p style="dir: ltr">`)
	fmt.Fprint(w, EscapingHtml[62])
	err = EscapingTemplates[30].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[63])
	fmt.Fprint(w, `<p style="border-left: 0; border-right: 1in">`)
	fmt.Fprint(w, EscapingHtml[64])
	err = EscapingTemplates[31].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[65])
	fmt.Fprint(w, `<p style="width: ZgotmplZ">`)
	fmt.Fprint(w, EscapingHtml[66])
	err = EscapingTemplates[32].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[67])
	fmt.Fprint(w, `<style>p { color: pink }</style>`)
	fmt.Fprint(w, EscapingHtml[68])
	err = EscapingTemplates[33].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[69])
	fmt.Fprint(w, `<p style="width: ZgotmplZ">`)
	fmt.Fprint(w, EscapingHtml[70])
	err = EscapingTemplates[34].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[71])
	fmt.Fprint(w, `<p style="ZgotmplZ: ...">`)
	fmt.Fprint(w, EscapingHtml[72])
	err = EscapingTemplates[35].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[73])
	fmt.Fprint(w, `<p style="background: url(/img?name=O%27Reilly%20Animal%281%29%3c2%3e.png)">`)
	fmt.Fprint(w, EscapingHtml[74])
	err = EscapingTemplates[36].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[75])
	fmt.Fprint(w, `<a style="background: url('#ZgotmplZ')">`)
	fmt.Fprint(w, EscapingHtml[76])
	err = EscapingTemplates[37].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[77])
	fmt.Fprint(w, `&iexcl;<b class="foo">Hello</b>, <textarea>O'World</textarea>!`)
	fmt.Fprint(w, EscapingHtml[78])
	err = EscapingTemplates[38].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[79])
	fmt.Fprint(w, `<div title="&iexcl;Hello, O&#39;World!">`)
	fmt.Fprint(w, EscapingHtml[80])
	err = EscapingTemplates[39].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[81])
	fmt.Fprint(w, `<button onclick="alert(&#34;&amp;iexcl;\u003cb class=\&#34;foo\&#34;\u003eHello\u003c/b\u003e, \u003ctextarea\u003eO&#39;World\u003c/textarea\u003e!&#34;)">`)
	fmt.Fprint(w, EscapingHtml[82])
	err = EscapingTemplates[40].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[83])
	fmt.Fprint(w, `<textarea>&iexcl;&lt;b class=&#34;foo&#34;&gt;Hello&lt;/b&gt;, &lt;textarea&gt;O&#39;World&lt;/textarea&gt;!</textarea>`)
	fmt.Fprint(w, EscapingHtml[84])
	err = EscapingTemplates[41].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[85])
	fmt.Fprint(w, `<input ZgotmplZ="doEvil()">`)
	fmt.Fprint(w, EscapingHtml[86])
	err = EscapingTemplates[42].Execute(w, data)
	handleEscapingError(err)
	fmt.Fprint(w, EscapingHtml[87])
}

func handleEscapingError(err error) {
	if err != nil {fmt.Println(err)}}