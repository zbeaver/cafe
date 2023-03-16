package render

import "github.com/zbeaver/cafe/pkg/vui"

type Head struct{}

func (r *Head) Render(n vui.INode) RenderFn {
	return RenderFn(func(slot string) string {
		return slot
	})
}
