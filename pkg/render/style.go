package render

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/zbeaver/cafe/pkg/vui"
	"golang.org/x/term"
)

var (
	painter         = (*styling)(nil)
	radio   float64 = 8.5
)

type styling struct {
	lipgloss.Style
	display string
	width   int
	height  int
}

type transformer func(vui.CSSStyleDecl) styling

func (s *styling) SetMaxSize(w int, h int) styling {
	s.height = h
	s.width = w
	return *s
}

// NewStyling return new styling instance
func NewStyling() styling {
	return styling{
		Style: lipgloss.NewStyle(),
	}
}

// Transform receive base styling (parent style)
// Compute and return new styling
// The following is attr applied from css:
//
// - margin
// - padding
// - border
// - background (inherit)
// - color (inherit)
// - display
//
// The following is rule applied specific cases:
//
// - width
//   always less than terminal width
//   updated when css.width less than base.width
//   base.width - base.padding
// - height
//   always less than terminal width
//   update if css.height less than base.height
//   base.height - base.padding
func TransformFrom(base styling) transformer {
	xW, xH, _ := term.GetSize(int(os.Stdout.Fd()))

	// mt, mr, mb, ml := base.GetMargin()
	// pt, pr, pb, pl := base.GetPadding()
	w := min(base.width, xW)
	h := min(base.height, xH)

	new := &styling{
		Style: lipgloss.
			NewStyle().
			Inherit(base.Style).
			MarginBackground(base.GetBackground()).
			Width(w),
		height: h,
		width:  w,
	}

	return transformer(func(css vui.CSSStyleDecl) styling {
		new.iWidth(css).
			iHeight(css).
			iMargin(css).
			iPadding(css).
			iAlign(css).
			iBackground(css).
			iForeground(css).
			iDisplay(css).
			iBorder(css)

		return *new
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
	if m == "" {
		return s
	}
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

	// s.width = s.width - s.GetMarginLeft() - s.GetMarginRight()
	// s.height = s.height - s.GetMarginBottom() - s.GetMarginTop()
	return s
}

func (s *styling) iPadding(css vui.CSSStyleDecl) *styling {
	m := css.GetPropertyValue("padding")
	if m == "" {
		return s
	}
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
	s.width = s.width - s.GetPaddingLeft() - s.GetPaddingRight()
	s.height = s.height - s.GetPaddingBottom() - s.GetPaddingTop()
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
		w = min(w, s.width)
		s.width = w
		s.Width(s.width)
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
		s.Width(s.GetWidth() - 4)
		s.Border(lipgloss.NormalBorder())
	}
	return s
}

func (s *styling) iDisplay(css vui.CSSStyleDecl) *styling {
	display := css.GetPropertyValue("display")
	s.display = display
	return s
}

func (s *styling) iAlign(css vui.CSSStyleDecl) *styling {
	return s
}
