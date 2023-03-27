package render

import (
	"github.com/zbeaver/cafe/pkg/vui"
)

type Br struct{}

func (r *Br) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Br) Render(n vui.INode, s styling, child string) string {
	return s.Render("\n\n")
}
