// This file is automatically generated by qtc from "page.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line examples/hu/templates/page.qtpl:1
package templates

//line examples/hu/templates/page.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line examples/hu/templates/page.qtpl:1
import "github.com/yanzay/teslo/templates"

//line examples/hu/templates/page.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line examples/hu/templates/page.qtpl:3
func StreamPage(qw422016 *qt422016.Writer, state State) {
	//line examples/hu/templates/page.qtpl:3
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body id="app">
    <form>
      <input type="text" name="repo"></input>
      <input type="submit" value="Add" id="add-repo">
    </form>
    `)
	//line examples/hu/templates/page.qtpl:13
	templates.StreamJS(qw422016)
	//line examples/hu/templates/page.qtpl:13
	qw422016.N().S(`
</body>
`)
//line examples/hu/templates/page.qtpl:15
}

//line examples/hu/templates/page.qtpl:15
func WritePage(qq422016 qtio422016.Writer, state State) {
	//line examples/hu/templates/page.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line examples/hu/templates/page.qtpl:15
	StreamPage(qw422016, state)
	//line examples/hu/templates/page.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line examples/hu/templates/page.qtpl:15
}

//line examples/hu/templates/page.qtpl:15
func Page(state State) string {
	//line examples/hu/templates/page.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line examples/hu/templates/page.qtpl:15
	WritePage(qb422016, state)
	//line examples/hu/templates/page.qtpl:15
	qs422016 := string(qb422016.B)
	//line examples/hu/templates/page.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line examples/hu/templates/page.qtpl:15
	return qs422016
//line examples/hu/templates/page.qtpl:15
}
