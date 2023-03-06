package layout

// + [AppWindow]* -----------------+
// |     context + window details  |
// | + Page -----------------------+
// | | Content          | Sidebar  |
// | |        + Dialog ---------+  |
// | |        |   Yes | Cancel  |  |
// | |        +-----------------+  |
// | |                  |          |
// | |                  |          |
// + AppStatus --------------------+
// |  <Focused:ComponentName>      +
// |-------------------------------+
// + AppCmd -----------------------+
// | >                             |
// +-------------------------------+

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zbeaver/cafe/pkg/vui"
)

type AppWindow struct {
	vui.Component
}

func (w *AppWindow) View() string {
	return "[DEBUG: APP_WINDOW]\n" + w.Component.SubView()
}

func (w *AppWindow) Init() tea.Cmd {
	t := &Test{}
	c := &Cmd{}
	w.SetChildren(t, c)
	return w.Component.Init()
}

type Test struct {
	vui.Component
}

func (t *Test) View() string {
	return "[DEBUG: TEST_COMPONENT]\n" + t.Component.SubView()
}
