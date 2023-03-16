package render

import "github.com/zbeaver/cafe/pkg/vui"

type Html struct{}

func (r *Html) Render(n vui.INode) RenderFn {
	return RenderFn(func(slot string) string {
		return slot
	})
}
