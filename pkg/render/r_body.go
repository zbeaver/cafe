package render

import "github.com/zbeaver/cafe/pkg/vui"

type Body struct{}

func (r *Body) Render(n vui.INode) RenderFn {
	return RenderFn(func(slot string) string {
		return slot
	})
}
