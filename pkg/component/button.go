package component

import "github.com/zbeaver/cafe/pkg/vui"

const TplButtonstring = `
	<button bg="{{ bg() }}" color="{{ color() }}">
	  {{ content }}
	</button>
`

type Button struct {
	vui.Component
}
