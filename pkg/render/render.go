package render

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/frankenbeanies/randhex"
	"github.com/zbeaver/cafe/pkg/vui"
)

type Registry map[uint32]Executor

type Executor interface {
	Render(vui.INode, styling, slots) string
	Style(styling, vui.INode) styling
}

type slots []string

type exec func(styling) string

const (
	TAG_BODY = iota + 1
	TAG_BR
	TAG_DIV
	TAG_HEAD
	TAG_HTML
	TAG_TEXT
	TAG_UNKNOWN
)

type engine struct {
	registry Registry
	ctx      context.Context
	doc      vui.Documentary
}

func NewEngine(ctx context.Context, doc vui.Documentary) *engine {
	reg := Registry{
		TAG_BODY:    (*Body)(nil),
		TAG_BR:      (*Unknown)(nil),
		TAG_DIV:     (*Div)(nil),
		TAG_HEAD:    (*Head)(nil),
		TAG_HTML:    (*Html)(nil),
		TAG_TEXT:    (*Text)(nil),
		TAG_UNKNOWN: (*Unknown)(nil),
	}

	return &engine{
		ctx:      ctx,
		doc:      doc,
		registry: reg,
	}
}

func (e *engine) Render() string {
	res := strings.Builder{}
	for _, c := range e.doc.ChildNodes() {
		res.WriteString(e.executor(c)(NewStyling()))
	}
	return res.String()
}

func debug(s styling, block string) string {
	w, h := lipgloss.Size(block)

	return s.Copy().Bold(true).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color(randhex.New().String())).
		Render(fmt.Sprintf("[%vx%v]", w, h))
}

func (e *engine) executor(node vui.INode) exec {
	var ex Executor

	switch node.(type) {
	case *vui.HtmlElm:
		ex = e.registry[TAG_HTML]
	case *vui.HeadElm:
		ex = e.registry[TAG_HEAD]
	case *vui.BrElm:
		ex = e.registry[TAG_BR]
	case *vui.BodyElm:
		ex = e.registry[TAG_BODY]
	case *vui.DivElm:
		ex = e.registry[TAG_DIV]
	case *vui.Text:
		ex = e.registry[TAG_TEXT]
	default:
		ex = e.registry[TAG_UNKNOWN]
	}

	return exec(func(base styling) string {
		var hook slots
		for _, c := range node.ChildNodes() {
			childView := strings.TrimSpace(e.executor(c)(ex.Style(base, c)))
			if childView != "" {
				hook = append(hook, childView)
			}
		}

		return strings.TrimSpace(ex.Render(node, base, hook))
	})
}
