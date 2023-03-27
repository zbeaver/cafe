package render

import (
	"github.com/zbeaver/cafe/pkg/vui"
)

type Div struct{}

func (r *Div) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Div) Render(n vui.INode, s styling, child string) string {
	return s.SetString(child).Render()
}
