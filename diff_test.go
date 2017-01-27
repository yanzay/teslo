package main

import "testing"

func TestDiff(t *testing.T) {
	initial := `
<ul>
  <li>first</li>
	<li>second</li>
</ul>
	`
	updated := `
<ul>
  <li>first</li>
	<li>second</li>
	<li>third</li>
</ul>
	`
	d := diff(initial, updated)
	if d.action != ActionAppendChild {
		t.Fail()
	}
	if d.content != "<li>third</li>" {
		t.Fail()
	}
}
