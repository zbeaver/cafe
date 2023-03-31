package render

import (
	"github.com/zbeaver/cafe/pkg/vui"
)

type Span struct{}

func (r *Span) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Span) Render(n vui.INode, s styling, child string) string {
	return s.SetString(child).Render()
}
