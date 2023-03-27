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

func (r *Text) Render(n vui.INode, s styling, child string) string {
	cont := strings.TrimSpace(n.NodeValue())
	if cont == "" || cont == "\n" {
		return ""
	}

	return cont
}
