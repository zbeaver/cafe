package render

import (
	"os"

	"github.com/zbeaver/cafe/pkg/vui"
	"golang.org/x/term"
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
	wt, ht, _ := term.GetSize(int(os.Stdout.Fd()))
	if wt > 0 {
		return s.MaxWidth(wt).Width(wt - 3).Height(ht - 2).Render(child...)
	}
	return s.Render(child...)
}
