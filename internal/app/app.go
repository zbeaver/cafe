package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zbeaver/cafe/internal/layout"
	"github.com/zbeaver/cafe/pkg/vui"
)

const TmpApp = `
<app>
  <tabs>
		<window title="personal" visibled>
		  Hello personal
		</window>
		<window title="zbeaver">
		  Hello</text>
		</window>
	</tabs>
	<status-bar>
	</status-bar>
	<command>
	</command>
<app>
`

type app struct {
	vui.Component
	Quitting bool
}

func NewApp() *app {
	return &app{
		Quitting: false,
	}
}

type History []*vui.Component

func (a *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle any top-level messages
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		}
	}
	_, cmd := a.Component.Update(msg)
	return a, cmd
}

func (a *app) View() string {
	return a.Component.SubView()
}

func (a *app) Init() tea.Cmd {
	win := &layout.AppWindow{}
	a.SetChildren(win)
	return a.Component.Init()
}
