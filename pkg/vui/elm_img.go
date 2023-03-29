package vui

import (
	"strconv"

	"golang.org/x/net/html"
)

type ImgElementary interface {
	Elementary
	Attributer
	Src() string
	Width() int
	Height() int
}

type ImgElm struct {
	*Elm
	src    string
	width  int
	height int
}

type ImgOpt func(*ImgElm)

var (
	_ Elementary = (*ImgElm)(nil)
)

func (e *ImgElm) ApplyAttr(attrs []html.Attribute) {
	for _, attr := range attrs {
		switch attr.Key {
		case "src":
			e.src = attr.Val
		case "width":
			e.width, _ = strconv.Atoi(attr.Val)
		case "height":
			e.height, _ = strconv.Atoi(attr.Val)
		}
	}
}

func (e *ImgElm) New(opts ...interface{}) Elementary {
	return &ImgElm{
		Elm: NewElm(opts...),
	}
}

func (e *ImgElm) Width() int {
	return e.width
}

func (e *ImgElm) Height() int {
	return e.height
}

func (e *ImgElm) Src() string {
	return e.src
}
