package render

import "github.com/zbeaver/cafe/pkg/vui"

type Div struct{}

func (r *Div) Render(n vui.INode) RenderFn {
	return RenderFn(func(slot string) string {
		return slot
	})
}
