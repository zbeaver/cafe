package render

import (
	"context"

	"github.com/ryboe/q"
	"github.com/zbeaver/cafe/pkg/vui"
)

type Render interface {
	Render(vui.INode) RenderFn
}

type RenderFn func(slot string) string

type Registry map[string]Render

type engine struct {
	registry Registry
	ctx      context.Context
	doc      vui.Documentary
}

func NewEngine(ctx context.Context, doc vui.Documentary) *engine {
	reg := Registry{
		"body": (*Body)(nil),
		"head": (*Head)(nil),
		"html": (*Html)(nil),
		"div":  (*Div)(nil),
		"text": (*Text)(nil),
	}

	return &engine{
		ctx:      ctx,
		doc:      doc,
		registry: reg,
	}
}

func (e *engine) Render() (result string) {
	for _, c := range e.doc.ChildNodes() {
		result += e.render(c)
	}
	return
}

func (e *engine) render(node vui.INode) string {
	var slot string
	for _, c := range node.ChildNodes() {
		slot += e.render(c)
	}

	switch node.(type) {
	case *vui.HtmlElm:
		fn := e.registry["html"].Render(node)
		return fn(slot)
	case *vui.HeadElm:
		fn := e.registry["head"].Render(node)
		return fn(slot)
	case *vui.BodyElm:
		fn := e.registry["body"].Render(node)
		return fn(slot)
	case *vui.DivElm:
		fn := e.registry["div"].Render(node)
		return fn(slot)
	case *vui.Text:
		fn := e.registry["text"].Render(node)
		return fn(slot)
	default:
		q.Q("COME HERE")
		return node.NodeValue()
	}
}
