// This file is automatically generated by qtc from "cart.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/cart.qtpl:1
package templates

//line templates/cart.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/cart.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/cart.qtpl:1
func StreamCartWidget(qw422016 *qt422016.Writer, items []*Item) {
	//line templates/cart.qtpl:1
	qw422016.N().S(`
  <div id="cart">
    <ul>
    `)
	//line templates/cart.qtpl:4
	for _, item := range items {
		//line templates/cart.qtpl:4
		qw422016.N().S(`
      <li>`)
		//line templates/cart.qtpl:5
		qw422016.E().S(item.Product.Name)
		//line templates/cart.qtpl:5
		qw422016.N().S(` - `)
		//line templates/cart.qtpl:5
		qw422016.N().D(item.Quantity)
		//line templates/cart.qtpl:5
		qw422016.N().S(`</li>
    `)
		//line templates/cart.qtpl:6
	}
	//line templates/cart.qtpl:6
	qw422016.N().S(`
    </ul>
  </div>
`)
//line templates/cart.qtpl:9
}

//line templates/cart.qtpl:9
func WriteCartWidget(qq422016 qtio422016.Writer, items []*Item) {
	//line templates/cart.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/cart.qtpl:9
	StreamCartWidget(qw422016, items)
	//line templates/cart.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line templates/cart.qtpl:9
}

//line templates/cart.qtpl:9
func CartWidget(items []*Item) string {
	//line templates/cart.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/cart.qtpl:9
	WriteCartWidget(qb422016, items)
	//line templates/cart.qtpl:9
	qs422016 := string(qb422016.B)
	//line templates/cart.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/cart.qtpl:9
	return qs422016
//line templates/cart.qtpl:9
}
