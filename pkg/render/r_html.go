package render

import (
	"os"

	"github.com/zbeaver/cafe/pkg/vui"
	"golang.org/x/term"
)

type Html struct{}

// calculate the term size and set to html
func (r *Html) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	wt, ht, _ := term.GetSize(int(os.Stdout.Fd()))
	if wt > 0 {
		base.SetMaxSize(wt, ht-2)
	}

	return TransformFrom(base)(elm.Style())
}

func (r *Html) Render(n vui.INode, s styling, child string) string {
	return s.SetString(child).Render()
}
