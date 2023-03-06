package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zbeaver/cafe/internal/app"
)

func main() {
	a := app.NewApp()
	p := tea.NewProgram(a, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program", err)
	}
}
