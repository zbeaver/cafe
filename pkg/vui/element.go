package vui

import "github.com/ryboe/q"

type Constructor interface {
	New(...interface{}) Elementary
}

type Elementary interface {
	INode
	SetInnerHTML(string)
}

type ElmOpt func(*Elm)

type Elm struct {
	*Node
	InnerHTML string
}

var (
	_ INode = (*Elm)(nil)
)

// @TODO: content must be decoder XML or HTML
func (el *Elm) SetInnerHTML(content string) {
	el.InnerHTML = content
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
			q.Q("Cannot cast", opt)
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
