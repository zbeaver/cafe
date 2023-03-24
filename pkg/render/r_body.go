package render

import (
	"github.com/zbeaver/cafe/pkg/vui"
)

type Body struct{}

func (r *Body) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}

	return TransformFrom(base)(elm.Style())
}

func (r *Body) Render(n vui.INode, s styling, child slots) string {
	if s.GetWidth() == 0 {
		s.Width(s.iMaxWidth)
	}

	if s.GetHeight() == 0 {
		s.Height(s.iMaxHeight)
	}

	return s.SetString(child...).Render()
}
