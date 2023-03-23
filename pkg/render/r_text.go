package render

import (
	"strings"

	"github.com/zbeaver/cafe/pkg/vui"
)

type Text struct{}

func (r *Text) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Text) Render(n vui.INode, s styling, child slots) string {
	if strings.TrimSpace(n.NodeValue()) == "" {
		return ""
	}
	return s.Render(strings.TrimSpace(n.NodeValue()))
}
