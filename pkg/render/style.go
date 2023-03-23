package render

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/zbeaver/cafe/pkg/vui"
)

var (
	painter = (*styling)(nil)
	radio   = 16
)

type styling struct {
	lipgloss.Style
	flex bool
}

type transformer func(vui.CSSStyleDecl) styling

func NewStyling() styling {
	return styling{
		Style: lipgloss.NewStyle(),
		flex:  false,
	}
}
func TransformFrom(base styling) transformer {
	new := &styling{
		Style: lipgloss.NewStyle().
			Background(base.GetBackground()).
			Foreground(base.GetForeground()).
			Align(base.GetAlign()),
		flex: base.flex,
	}

	return transformer(func(css vui.CSSStyleDecl) styling {
		new.iMargin(css).
			iPadding(css).
			iAlign(css).
			iBackground(css).
			iForeground(css).
			iDisplay(css).
			iWidth(css).
			iHeight(css)
		return *new
	})
}

// func (s *style) toLipgloss() {
// 	for s.css.GetPropertyValue("padding")
// }

func strconvInt(str string) (int, error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(reg.ReplaceAllString(str, ""))
	return int(n / radio), err
}

func parseMeasureString(m string) []int {
	p := make([]int, 0)

	for _, f := range strings.SplitN(m, " ", 4) {
		val, _ := strconvInt(f)
		p = append(p, val)
	}

	return p
}

func (s *styling) iMargin(css vui.CSSStyleDecl) *styling {
	m := css.GetPropertyValue("margin")
	ret := parseMeasureString(m)

	if str := css.GetPropertyValue("margin-top"); str != "" {
		ret[0], _ = strconvInt(str)
	}

	if str := css.GetPropertyValue("margin-right"); str != "" {
		ret[1], _ = strconvInt(str)
	}

	if str := css.GetPropertyValue("margin-bottom"); str != "" {
		ret[2], _ = strconvInt(str)
	}

	if str := css.GetPropertyValue("margin-left"); str != "" {
		ret[3], _ = strconvInt(str)
	}

	s.Margin(ret...)
	return s
}

func (s *styling) iPadding(css vui.CSSStyleDecl) *styling {
	m := css.GetPropertyValue("padding")
	ret := parseMeasureString(m)

	if str := css.GetPropertyValue("padding-top"); str != "" {
		ret[0], _ = strconvInt(str)
	}

	if str := css.GetPropertyValue("padding-right"); str != "" {
		ret[1], _ = strconvInt(str)
	}

	if str := css.GetPropertyValue("padding-bottom"); str != "" {
		ret[2], _ = strconvInt(str)
	}

	if str := css.GetPropertyValue("padding-left"); str != "" {
		ret[3], _ = strconvInt(str)
	}

	s.Padding(ret...)
	return s
}

func (s *styling) iBackground(css vui.CSSStyleDecl) *styling {
	if str := css.GetPropertyValue("background"); str != "" {
		s.Background(lipgloss.Color(str))
	}
	if str := css.GetPropertyValue("background-color"); str != "" {
		s.Background(lipgloss.Color(str))
	}
	return s
}

func (s *styling) iForeground(css vui.CSSStyleDecl) *styling {
	if str := css.GetPropertyValue("color"); str != "" {
		s.Foreground(lipgloss.Color(str))
	}
	return s
}

func (s *styling) iWidth(css vui.CSSStyleDecl) *styling {
	if str := css.GetPropertyValue("width"); str != "" {
		w, _ := strconvInt(str)
		s.Width(w)
	}
	return s
}

func (s *styling) iHeight(css vui.CSSStyleDecl) *styling {
	if str := css.GetPropertyValue("height"); str != "" {
		h, _ := strconvInt(str)
		s.Height(h)
	}
	return s
}

func (s *styling) bBorder(css vui.CSSStyleDecl) *styling {
	return s
}

func (s *styling) iDisplay(css vui.CSSStyleDecl) *styling {
	str := css.GetPropertyValue("display")
	switch str {
	case "flex":
		s.flex = true
	case "block":
		// set horizonal join
	default:
		// use default
	}

	return s
}

func (s *styling) iAlign(css vui.CSSStyleDecl) *styling {
	return s
}
