package render

import "github.com/zbeaver/cafe/pkg/vui"

type Text struct{}

func (r *Text) Render(n vui.INode) RenderFn {
	return RenderFn(func(slot string) string {
		return n.NodeValue()
	})
}
