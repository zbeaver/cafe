package vui

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/ryboe/q"
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

	var (
		count = 1
		deep  = 0
	)

	d := xml.NewDecoder(bytes.NewReader(vl.template))
	for {
		t, err := d.Token()
		q.Q(">>>>>>>>>>", t)
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		// When token is StartElement
		// - All content inside StartElement become $slots
		// - Find responding component in .components attr
		// - Create new empty component and pass $slots to inside
		// - Create new loader and run loader.Load
		// When token is ChartData
		// - Create PlainComponent for this case
		// - Append to $slot related
		case xml.StartElement:
			q.Q("start", tt.Name.Local)
			if count == 1 {
				count++
				continue
			}
			deep++
			for key, val := range vl.components {
				if key == tt.Name.Local {
					nested := val()
					vl.root.AppendChild(nested)
				} else {
					return fmt.Errorf("component %s not found", key)
				}
			}
		//
		case xml.EndElement:
			if count == 1 {
				continue
			}
			q.Q("end", tt.Name.Local)
			deep--
			count--
		case xml.CharData:
			q.Q("content", string(tt))
		}
	}

	return nil
}
