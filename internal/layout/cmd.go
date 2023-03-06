package layout

import "github.com/zbeaver/cafe/pkg/vui"

type Cmd struct {
	vui.Component
}

func (c *Cmd) View() string {
	return "[DEBUG: COMMAND]\n"
}
