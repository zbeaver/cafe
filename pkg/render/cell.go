package render

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Tissue
// +------------------------------------+
// |+---------------------------------+ |
// || +------+   +------+    Cell(s)  | |
// || |Cell  |   | Cell |             | |
// || +------+   +------+             | |
// |+---------------------------------+ |
// |+---------------------------------+ |
// ||   +------+            Cell(s)   | |
// ||   |Cell  |                      | |
// ||   +------+                      | |
// |+---------------------------------+ |
// |                                    |
// +------------------------------------+

type grapher interface {
	Size() (int, int)
	Position() (int, int)
	Render() string
}

type position struct {
	x int
	y int
}

const (
	DISPLAY_BLOCK = iota + 1
	DISPLAY_CONTENT
	DISPLAY_INLINE
	DISPLAY_INLINE_BLOCK
	DISPLAY_FLEX
	DISPLAY_NONE
)

type cells []cell

type cell struct {
	raw string
	w   int
	h   int
	d   int
	position
}

type CellOpt func(*cell)

func WithPosition(x int, y int) func(*cell) {
	return func(c *cell) {
		c.position = position{
			x: x,
			y: y,
		}
	}
}

func WithDisplay(d string) func(*cell) {
	return func(c *cell) {
		switch d {
		case "block":
			c.d = DISPLAY_BLOCK
		case "content":
			c.d = DISPLAY_CONTENT
		case "inline":
			c.d = DISPLAY_INLINE
		case "inline-block":
			c.d = DISPLAY_INLINE_BLOCK
		case "flex":
			c.d = DISPLAY_FLEX
		case "none":
			c.d = DISPLAY_NONE
		default:
			c.d = DISPLAY_BLOCK
		}
	}
}

func NewCell(raw string, opts ...CellOpt) cell {
	cell := cell{
		raw: raw,
	}

	w, h := lipgloss.Size(raw)
	cell.w = w
	cell.h = h

	for _, opt := range opts {
		opt(&cell)
	}

	return cell
}

func (c cell) Render() string {
	return c.raw
}

func (c cell) Size() (int, int) {
	return c.w, c.h
}

func (c cell) Position() (int, int) {
	return c.position.x, c.position.y
}

type tissue struct {
	cells []cells
	s     *styling
	w     int
	h     int
	position
}

func FillBg(color lipgloss.TerminalColor, src string, w, h int) string {
	style := lipgloss.NewStyle().Background(color)
	ws, hs := lipgloss.Size(src)
	lines := strings.Split(src, "\n")

	if w > ws {
		for idx := range lines {
			lines[idx] += style.Copy().UnsetWidth().SetString(strings.Repeat(" ", w-ws)).Render()
		}
	}

	if h > hs {
		for i := 0; i < h-hs; i++ {
			lines = append(lines, style.Copy().UnsetWidth().SetString(strings.Repeat(" ", max(ws, w))).Render())
		}
	}

	return strings.Join(lines, "\n")
}

// Tissue is matrix of cell: [][]cell
// Tissue receive list cells and compute to matrix cells adapt these rules:
// - If cell is display block - It must display reside inside one row
// - If cell is display-block, it able to append to previous row
// - If cell is display-block the total width of all cells in a row not great than iMaxWidth
func NewTissue(cs cells, s *styling) *tissue {
	t := &tissue{s: s}
	// This value use to determine condition max width matched
	widthLine := 0
	for _, c := range cs {
		switch c.d {
		// Block mean
		// It starts on a new line, and takes up the whole width
		case DISPLAY_BLOCK:
			t.cells = append(t.cells, cells{c}, make(cells, 0))
			// reset widthLine
			widthLine = 0

		// Displays an element as an inline-level block container.
		// The element itself is formatted as an inline element,
		// but you can apply height and width values
		//
		// When cell is inline block keep continue append to row
		case DISPLAY_INLINE_BLOCK:
			if len(t.cells) == 0 {
				w, _ := lipgloss.Size(c.Render())
				t.cells = append(t.cells, cells{c})
				widthLine = w
				continue
			}

			w, _ := lipgloss.Size(c.Render())
			widthLine += w
			// If total width of cells adapt condition less than maxwidth
			if widthLine <= s.width {
				curr := len(t.cells) - 1
				t.cells[curr] = append(t.cells[curr], c)
			} else {
				// When total width great than maxwidth,
				// break it down to new row
				widthLine = w
				t.cells = append(t.cells, cells{c})
			}

		// Inline mean
		// Displays an element as an inline element (like <span>).
		// Any height and width properties will have no effect
		case DISPLAY_INLINE:
			// @TODO leave it now

		// flex mean
		// Displays an element as a block-level flex container
		case DISPLAY_FLEX:
			// @TODO leave it now

		case DISPLAY_NONE:
			// @TODO leave it now

		default:
			// same as display block
			t.cells = append(t.cells, cells{c}, make(cells, 0))
			widthLine = 0
		}
	}

	if len(t.cells[len(t.cells)-1]) == 0 {
		// remove last item
		t.cells = t.cells[:len(t.cells)-1]
	}
	return t
}

func (t *tissue) Size() (int, int) {
	return t.w, t.h
}

func (t *tissue) Position() (int, int) {
	return t.position.x, t.position.y
}

// Render return string from tissue.cells
// Firstly,
//   - Loop each row at tissue.cells
//   - render each cell
//   - use JoinHorizontal to merge result to []string _rows
// Seconds
//   - Loop each _rows
//   - use JoinVertical to merge result from _rows
func (t *tissue) Render() (result string) {
	// filter the empty tissue
	if len(t.cells) == 0 {
		t.w = 0
		t.h = 0
		return ""
	}

	rows := make([]string, 0)
	for _, line := range t.cells {
		if len(line) > 0 {
			var maxHeight int
			cols := make([]string, 0)
			cells := make([]string, 0)
			// Draw content first
			for _, c := range line {
				maxHeight = max(c.h, maxHeight)
				cols = append(cols, c.Render())
			}

			// fill background by height
			for _, c := range cols {
				cells = append(cells, FillBg(
					t.s.GetBackground(),
					c,
					0,
					maxHeight,
				))
			}

			// row := lipgloss.JoinHorizontal(lipgloss.Top, cells...)
			// fill background by weight
			row := FillBg(
				t.s.GetBackground(),
				lipgloss.JoinHorizontal(lipgloss.Top, cells...),
				t.s.width,
				0,
			)

			// append parent
			rows = append(rows, row)
		}
	}
	result = lipgloss.JoinVertical(lipgloss.Top, rows...)
	// store size of block
	t.w, t.h = lipgloss.Size(result)
	return result
}
