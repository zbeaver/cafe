package vui

import (
	"strings"
)

type Constructor interface {
	New(...interface{}) Elementary
}

type Elementary interface {
	INode
	SetInnerHTML(string)
	InnerHTML() string
	Style() CSSStyleDecl
	Class() string
	ClassList() []string
	SetClass(...string)
	SetId(string)
	Id() string
}

type ElmOpt func(*Elm)

func WithStyle(text string) ElmOpt {
	style := NewCSSStyleDecl(text)
	return ElmOpt(func(el *Elm) {
		el.style = style
	})
}

func WithId(id string) ElmOpt {
	return ElmOpt(func(el *Elm) {
		el.id = id
	})
}

func WithClass(classes ...string) ElmOpt {
	return ElmOpt(func(el *Elm) {
		el.SetClass(classes...)
	})
}

type Elm struct {
	*Node
	innerHTML string
	style     CSSStyleDecl
	id        string
	className string
}

var (
	_ INode = (*Elm)(nil)
)

// type InlineStyle struct {}

// type Style map[string]string

// @TODO: content must be decoder XML or HTML
func (el *Elm) SetInnerHTML(content string) {
	el.innerHTML = content
}

func (el *Elm) InnerHTML() string {
	return el.innerHTML
}

func (el *Elm) Style() CSSStyleDecl {
	return el.style
}

func (el *Elm) Class() string {
	return el.className
}

func (el *Elm) ClassList() []string {
	return strings.Split(el.className, " ")
}

func (el *Elm) SetClass(classes ...string) {
	el.className = strings.Join(classes, " ")
}

func (el *Elm) SetId(id string) {
	el.id = id
}

func (el *Elm) Id() string {
	return el.id
}

func NewElm(opts ...interface{}) *Elm {
	elmOpt := make([]ElmOpt, 0, 10)
	nodeOpt := make([]NodeOpt, 0, 10)

	for _, opt := range opts {
		switch opt.(type) {
		case ElmOpt:
			o, _ := opt.(ElmOpt)
			elmOpt = append(elmOpt, o)
		case NodeOpt:
			o, _ := opt.(NodeOpt)
			nodeOpt = append(nodeOpt, o)
		default:
		}
	}

	el := &Elm{
		Node: NewNode(nodeOpt...),
	}

	for _, opt := range elmOpt {
		opt(el)
	}

	return el
}
