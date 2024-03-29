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
	Render(vui.INode, styling, string) string
	Style(styling, vui.INode) styling
}

type exec func(styling) string

const (
	TAG_BODY = iota + 1
	TAG_BR
	TAG_BUTTON
	TAG_DIV
	TAG_FIELDSET
	TAG_FORM
	TAG_HEAD
	TAG_HEADING
	TAG_HR
	TAG_HTML
	TAG_IMG
	TAG_INPUT
	TAG_LABEL
	TAG_LEGEND
	TAG_OPTGROUP
	TAG_OPTION
	TAG_PARAGRAPH
	TAG_SELECT
	TAG_SPAN
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
		TAG_BODY:      (*Body)(nil),
		TAG_BR:        (*Br)(nil),
		TAG_DIV:       (*Div)(nil),
		TAG_FIELDSET:  (*FieldSet)(nil),
		TAG_FORM:      (*Form)(nil),
		TAG_HEAD:      (*Head)(nil),
		TAG_HEADING:   (*Heading)(nil),
		TAG_HR:        (*Hr)(nil),
		TAG_HTML:      (*Html)(nil),
		TAG_IMG:       (*Img)(nil),
		TAG_INPUT:     (*Input)(nil),
		TAG_LABEL:     (*Label)(nil),
		TAG_LEGEND:    (*Legend)(nil),
		TAG_OPTGROUP:  (*OptGroup)(nil),
		TAG_OPTION:    (*Option)(nil),
		TAG_PARAGRAPH: (*Paragraph)(nil),
		TAG_SELECT:    (*Select)(nil),
		TAG_SPAN:      (*Span)(nil),
		TAG_TEXT:      (*Text)(nil),
		TAG_UNKNOWN:   (*Unknown)(nil),
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
	case *vui.BodyElm:
		ex = e.registry[TAG_BODY]
	case *vui.BrElm:
		ex = e.registry[TAG_BR]
	case *vui.DivElm:
		ex = e.registry[TAG_DIV]
	case *vui.FieldSetElm:
		ex = e.registry[TAG_FIELDSET]
	case *vui.FormElm:
		ex = e.registry[TAG_FORM]
	case *vui.HeadingElm:
		ex = e.registry[TAG_HEADING]
	case *vui.HeadElm:
		ex = e.registry[TAG_HEAD]
	case *vui.HrElm:
		ex = e.registry[TAG_HR]
	case *vui.HtmlElm:
		ex = e.registry[TAG_HTML]
	case *vui.ImgElm:
		ex = e.registry[TAG_IMG]
	case *vui.InputElm:
		ex = e.registry[TAG_INPUT]
	case *vui.LabelElm:
		ex = e.registry[TAG_LABEL]
	case *vui.LegendElm:
		ex = e.registry[TAG_LEGEND]
	case *vui.OptGroupElm:
		ex = e.registry[TAG_OPTGROUP]
	case *vui.OptionElm:
		ex = e.registry[TAG_OPTION]
	case *vui.ParagraphElm:
		ex = e.registry[TAG_PARAGRAPH]
	case *vui.SelectElm:
		ex = e.registry[TAG_SELECT]
	case *vui.SpanElm:
		ex = e.registry[TAG_SPAN]
	case *vui.Text:
		ex = e.registry[TAG_TEXT]
	default:
		ex = e.registry[TAG_UNKNOWN]
	}

	return exec(func(base styling) string {
		cs := make(cells, 0)

		for _, c := range node.ChildNodes() {
			// rescurive call
			styling := ex.Style(base, c)
			subview := strings.TrimSpace(e.executor(c)(styling))
			if subview != "" {
				cs = append(cs, NewCell(subview, WithDisplay(styling.display)))
			}
		}

		if len(cs) > 0 {
			t := NewTissue(cs, &base)
			return strings.TrimSpace(ex.Render(node, base, t.Render()))
		} else {
			return strings.TrimSpace(ex.Render(node, base, ""))
		}
	})
}
