package render

import (
	"strings"

	"github.com/zbeaver/cafe/pkg/vui"
)

type Text struct{}

func (r *Text) Render(n vui.INode) RenderFn {
	return RenderFn(func(slot string) string {
		return strings.Trim(n.NodeValue(), "\n\t")
	})
}
