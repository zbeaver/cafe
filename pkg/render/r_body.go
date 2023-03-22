package render

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/zbeaver/cafe/pkg/vui"
)

type Body struct{}

func (r *Body) Render(n vui.INode) RenderFn {
	elm, _ := n.(vui.Elementary)
	p, _ := strconv.Atoi(elm.Style().GetPropertyValue("padding"))
	div := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(elm.Style().GetPropertyValue("foreground-color"))).
		Background(lipgloss.Color(elm.Style().GetPropertyValue("background-color"))).
		Padding(p)

	return RenderFn(func(slot string) string {
		return div.Render(slot)
	})
}
