package vui

import (
	"encoding/xml"

	"github.com/ryboe/q"
)

type Loader interface {
	Load() Elm
}

type VuiLoader struct {
	raw        []byte
	components map[string]*Component
}

func NewLoader(r []byte) Loader {
	vl := &VuiLoader{
		raw: r,
	}
	return vl
}

func (vl *VuiLoader) Load() Elm {
	c := &Component{}
	if err := xml.Unmarshal(vl.raw, c); err != nil {
		q.Q(err)
	}

	q.Q(c)
	return c
}
