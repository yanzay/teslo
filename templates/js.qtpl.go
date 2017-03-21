// This file is automatically generated by qtc from "js.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/js.qtpl:1
package templates

//line templates/js.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/js.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/js.qtpl:1
func StreamJS(qw422016 *qt422016.Writer) {
	//line templates/js.qtpl:1
	qw422016.N().S(`
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
<script>
window.addEventListener("load", function(evt) {
  var ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = function(e) {
    console.log("OPEN");
  };
  ws.onclose = function(e) {
    console.log("CLOSE");
  };
  ws.onmessage = function(e) {
    var message = JSON.parse(e.data);
    var el = document.getElementById(message.id)
    el.outerHTML = message.content;
  };
  ws.onerror = function(e) {
    console.log("Error: ", e.data);
  };
  var app = document.getElementById("app")

  var clickHandler = function(e) {
    console.log(e.target.id);
    var parentIds = $(e.target).parents().map(function(i, el) {return el.id;}).toArray().filter(function(id) {return id !== ""});
    console.log(parentIds);
    if (e.target.id) {
      ws.send(JSON.stringify({event: "click", id: e.target.id, parents: parentIds}));
    }
  };
  var submitHandler = function(e) {
    e.preventDefault();
    var form = $(e.target).serialize();
    console.log(form);
    form.event = "submit";
    form.id = e.target.id;
    console.log(form);
    ws.send(JSON.stringify(form));
  };
  app.addEventListener("click", clickHandler);
  app.addEventListener("submit", submitHandler);
});
</script>
`)
//line templates/js.qtpl:43
}

//line templates/js.qtpl:43
func WriteJS(qq422016 qtio422016.Writer) {
	//line templates/js.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/js.qtpl:43
	StreamJS(qw422016)
	//line templates/js.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line templates/js.qtpl:43
}

//line templates/js.qtpl:43
func JS() string {
	//line templates/js.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/js.qtpl:43
	WriteJS(qb422016)
	//line templates/js.qtpl:43
	qs422016 := string(qb422016.B)
	//line templates/js.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/js.qtpl:43
	return qs422016
//line templates/js.qtpl:43
}
