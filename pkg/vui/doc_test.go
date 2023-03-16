package vui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test default element registered
func TestDoc_Children(t *testing.T) {
	tags := []string{
		"div",
	}

	for _, tag := range tags {
		doc := NewDocument()
		el, _ := doc.CreateElement(tag)
		if err := doc.AppendChild(el); err != nil {
			panic(err)
		}
	}
}

func TestDoc_Factory(t *testing.T) {
	as := assert.New(t)
	f := &elmFactory{}
	f_div, _ := f.Factory("div")
	el, err := f_div()
	as.Equal(err, nil)
	if _, ok := el.(*DivElm); !ok {
		panic("elem not is DivElm")
	}
}
