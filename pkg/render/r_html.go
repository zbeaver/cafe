package render

import (
	"github.com/zbeaver/cafe/pkg/vui"
)

type Html struct{}

func (r *Html) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Html) Render(n vui.INode, s styling, child slots) string {
	return s.Render(child...)
}
