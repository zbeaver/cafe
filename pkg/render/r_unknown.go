package render

import (
	"github.com/zbeaver/cafe/pkg/vui"
)

type Unknown struct{}

func (r *Unknown) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Unknown) Render(n vui.INode, s styling, child string) string {
	return "!!unknown!!"
}
