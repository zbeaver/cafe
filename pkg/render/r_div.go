package render

import (
	"github.com/76creates/stickers"
	"github.com/charmbracelet/lipgloss"
	"github.com/zbeaver/cafe/pkg/vui"
)

type Div struct{}

func (r *Div) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Div) Render(n vui.INode, s styling, child slots) string {

	if s.flex {
		flexbox := stickers.NewFlexBox(0, 0)
		var contents []*stickers.FlexBoxCell
		for _, c := range child {
			contents = append(
				contents,
				stickers.NewFlexBoxCell(1, 6).SetStyle(s.Style).SetContent(c),
			)
		}
		rows := []*stickers.FlexBoxRow{
			flexbox.NewRow().AddCells(contents),
		}
		return flexbox.AddRows(rows).Render()
	}

	if s.inlineBlock {
		return s.
			Copy().
			MaxWidth(s.iMaxWidth).
			SetString(lipgloss.JoinHorizontal(lipgloss.Left, child...)).
			Render()
	} else {
		ret := lipgloss.JoinVertical(lipgloss.Top, child...)
		if s.GetWidth() > 0 {
			return s.Copy().SetString(ret).Render()
		} else {
			return s.Copy().Width(s.iMaxWidth).SetString(ret).Render()
		}
	}
}
