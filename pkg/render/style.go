package render

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/zbeaver/cafe/pkg/vui"
)

var (
	painter         = (*styling)(nil)
	radio   float64 = 8.5
)

type styling struct {
	lipgloss.Style
	flex        bool
	inlineBlock bool
	iMaxWidth   int
	iMaxHeight  int
}

type transformer func(vui.CSSStyleDecl) styling

func (s *styling) SetMaxSize(w int, h int) styling {
	s.iMaxHeight = h
	s.iMaxWidth = w
	return *s
}

func NewStyling() styling {
	return styling{
		Style: lipgloss.NewStyle(),
	}
}

func TransformFrom(base styling) transformer {
	new := styling{
		Style:       lipgloss.NewStyle().Inherit(base.Style),
		flex:        base.flex,
		iMaxHeight:  base.iMaxHeight,
		iMaxWidth:   base.iMaxWidth,
		inlineBlock: base.inlineBlock,
	}

	return transformer(func(css vui.CSSStyleDecl) styling {
		new.iMargin(css).
			iPadding(css).
			iAlign(css).
			iBackground(css).
			iForeground(css).
			iDisplay(css).
			iWidth(css).
			iHeight(css).
			iBorder(css)
		return new
	})
}

func strconvInt(str string) (int, error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(reg.ReplaceAllString(str, ""))
	return int(math.Round(float64(n) / radio)), err
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
	var err error
	if str := css.GetPropertyValue("margin-top"); str != "" {
		ret[0], err = strconvInt(str)
	}

	if str := css.GetPropertyValue("margin-right"); str != "" {
		ret[1], err = strconvInt(str)
	}

	if str := css.GetPropertyValue("margin-bottom"); str != "" {
		ret[2], err = strconvInt(str)
	}

	if str := css.GetPropertyValue("margin-left"); str != "" {
		ret[3], err = strconvInt(str)
	}

	if err != nil {
		log.Fatal(err)
	}
	s.Margin(ret...)

	// s.Width(s.GetWidth() - s.GetMarginLeft() - s.GetMarginRight())
	return s
}

func (s *styling) iPadding(css vui.CSSStyleDecl) *styling {
	m := css.GetPropertyValue("padding")
	ret := parseMeasureString(m)
	var err error
	if str := css.GetPropertyValue("padding-top"); str != "" {
		ret[0], err = strconvInt(str)
	}

	if str := css.GetPropertyValue("padding-right"); str != "" {
		ret[1], err = strconvInt(str)
	}

	if str := css.GetPropertyValue("padding-bottom"); str != "" {
		ret[2], err = strconvInt(str)
	}

	if str := css.GetPropertyValue("padding-left"); str != "" {
		ret[3], err = strconvInt(str)
	}

	if err != nil {
		log.Fatal(err)
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
		w, err := strconvInt(str)
		if err != nil {
			log.Fatal(err)
		}
		s.Width(w)
	}
	return s
}

func (s *styling) iHeight(css vui.CSSStyleDecl) *styling {
	if str := css.GetPropertyValue("height"); str != "" {
		h, err := strconvInt(str)
		if err != nil {
			log.Fatal(err)
		}
		s.Height(h)
	}
	return s
}

func (s *styling) iBorder(css vui.CSSStyleDecl) *styling {
	if str := css.GetPropertyValue("border"); str != "" {
		// s.MaxHeight = s.MaxHeight - 30
		// s.MaxWidth = s.MaxWidth - 30
		// s.Width(s.MaxWidth).Height(s.MaxHeight).Border(lipgloss.NormalBorder())
		s.Border(lipgloss.NormalBorder())
	}
	return s
}

func (s *styling) iDisplay(css vui.CSSStyleDecl) *styling {
	str := css.GetPropertyValue("display")
	switch str {
	case "flex":
		s.flex = true
	case "inline-block":
		s.inlineBlock = true
	default:
	}

	return s
}

func (s *styling) iAlign(css vui.CSSStyleDecl) *styling {
	return s
}
