package vui

import (
	"testing"
)

type A struct {
	vui.Component
}

var (
	_ vui.Template = (*A)(nil)
)

func (a *A) Template() []byte {
	return []byte(`
		<div :inline="true">
			<div>hello,</div>
			<div>world</div>
		</div>
		<div :inline="false">
			<div>hello</div>
			<div>world</div>
		</div>
	`)
}

func TestDiv(t *testing.T) {
	// c := Div{}
	a := vui.A{}
}
