package vui

import (
	"bytes"
	"encoding/xml"
	"io"
)

type ElmFn func() Loadable
type Registry map[string]ElmFn

type Template interface {
	Template() []byte
	Components() Registry
}

type Loadable interface {
	Elm
	Template
}

type Loader interface {
	Load() error
}

type loader struct {
	template   []byte
	root       Elm
	components Registry
}

func NewLoader(el Loadable) Loader {
	return &loader{
		template:   el.Template(),
		root:       el,
		components: el.Components(),
	}
}

// Load return component initialed
func (vl *loader) Load() error {
	if len(vl.template) == 0 {
		return nil
	}

	d := xml.NewDecoder(bytes.NewReader(vl.template))
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.StartElement:
			for key, val := range vl.components {
				if key == tt.Name.Local {
					nested := val()
					vl.root.AppendChild(nested)
					// ld := NewLoader(nested)
					// ld.Load()
				} else {
					// Ignore
					// panic("component not found")
				}
			}
		}
	}

	return nil
}
