package vui

import (
	"fmt"
)

type Documentary interface {
	INode
	CreateElement(string, ...interface{}) (Elementary, error)
	CreateText(string) (Texter, error)
}

type Document struct {
	*Node
	ActiveElement     Elm
	AlinkColor        HexColor
	Body              Elm
	ChildElementCount int
	elmFactory        ElmFactory
}

// Factory method use to create new element
type ElmFactory interface {
	Factory(tag string) (elmFacMethod, error)
}

type elmFactory struct{}
type elmFacMethod func(...interface{}) (Elementary, error)

// Registry use to register default elements
type ElmRegistry map[string]Elementary

var (
	registry = ElmRegistry{
		"html": (*HtmlElm)(nil),
		"body": (*BodyElm)(nil),
		"head": (*HeadElm)(nil),
		"div":  (*DivElm)(nil),
	}
	_ Documentary = (*Document)(nil)
)

// CustomElmRegistry is runtime regiter custom elemenent
type CustomElemRegistry map[string]Elementary

func (f *elmFactory) Factory(tag string) (elmFacMethod, error) {
	xel := registry[tag]
	if xel == nil {
		return nil, fmt.Errorf("The element tag name `%v` not found", tag)
	}

	return func(opts ...interface{}) (Elementary, error) {
		elm, ok := xel.(Constructor)
		if !ok {
			return nil, fmt.Errorf("the element tag name `%v` missing new func", tag)
		}
		e := elm.New(opts...)
		return e, nil
	}, nil
}

func NewDocument() *Document {
	return &Document{
		elmFactory: &elmFactory{},
		Node:       NewNode(),
	}
}

func (c *Document) CreateText(val string) (Texter, error) {
	var t *Text
	return t.New(val), nil
}

func (d *Document) CreateElement(tag string, opts ...interface{}) (Elementary, error) {
	f, err := d.elmFactory.Factory(tag)
	if err != nil {
		panic(err)
	}
	elm, err := f(opts...)
	if err != nil {
		panic(err)
		// return nil, err
	}
	elm.SetOwnerDocument(d)
	return elm, nil
}
