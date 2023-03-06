package vui

import tea "github.com/charmbracelet/bubbletea"

type Cmd interface {
}

type Keystoke string

type CmdName string

type CmdHandler func(tea.Model, tea.Cmd)

type ListCmd struct {
	key    Keystoke
	cmd    CmdName
	alias  string
	handle CmdHandler
}
