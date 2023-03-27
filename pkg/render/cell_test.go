package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell__Simple(t *testing.T) {
	cases := []struct {
		raw      string
		expected func(*assert.Assertions, *cell)
		opts     []CellOpt
	}{
		{
			raw: "this is 2 line content\nhello world",
			expected: func(as *assert.Assertions, c *cell) {
				as.Equal(2, c.h)
				as.Equal(22, c.w)
			},
		},
		{
			raw: "3 line\n1234567890\n1234567890",
			expected: func(as *assert.Assertions, c *cell) {
				as.Equal(3, c.h)
				as.Equal(10, c.w)
			},
		},
		{
			raw: "testing",
			opts: []CellOpt{
				WithDisplay("block"),
				WithPosition(10, 10),
			},
			expected: func(as *assert.Assertions, c *cell) {
				as.Equal(DISPLAY_BLOCK, c.d)
				as.Equal(10, c.position.x)
				as.Equal(10, c.position.y)
			},
		},
	}

	// Run test cases
	as := assert.New(t)
	for _, tc := range cases {
		c := NewCell(tc.raw, tc.opts...)
		tc.expected(as, &c)
	}
}

func TestTissue__New(t *testing.T) {
	cases := []struct {
		cells    func() cells
		expected func(*assert.Assertions, *tissue)
	}{
		// Default case all NewCell is block
		{
			cells: func() cells {
				return cells{
					NewCell("one"),
					NewCell("two"),
					NewCell("three"),
					NewCell("fourth"),
				}
			},
			expected: func(as *assert.Assertions, t *tissue) {
				as.Equal("one   \ntwo   \nthree \nfourth", t.Render())
				as.Equal(4, t.h)
				as.Equal(6, t.w)
			},
		},
		// Include inline-block
		{
			cells: func() cells {
				return cells{
					NewCell("one"),
					NewCell("two", WithDisplay("inline-block")),
					NewCell("three", WithDisplay("inline-block")),
					NewCell("fourth", WithDisplay("block")),
				}
			},
			expected: func(as *assert.Assertions, t *tissue) {
				as.Equal("one     \ntwothree\nfourth  ", t.Render())
				as.Equal(3, t.h)
				as.Equal(8, t.w)
			},
		},
	}

	// Run test cases
	as := assert.New(t)
	for _, tc := range cases {
		s := NewStyling()
		tc.expected(as, NewTissue(tc.cells(), &s))
	}
}
