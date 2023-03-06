package vui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Elm interface {
	tea.Model
	SetParent(Elm) error
	SetChildren(...Elm) error
}

type Focusable interface {
	Focused() error
	Unfocus() error
	IsFocused() bool
}

type Switcher interface {
	Show() Elm
	Hide() Elm
}

type Component struct {
	// determine for component have been joined to tree-struct
	initialized bool
	// if the parent = nil and initialized = true mean
	// it's root component
	parent   Elm
	children []Elm

	// active state of component
	focused bool

	// when visible is true, the component have been showed
	visible bool

	// template
	template []byte
}

var (
	_ Focusable = (*Component)(nil)
)

// NewRootComponent return a component root
func NewRootComponent() *Component {
	return &Component{
		parent:      nil,
		focused:     true,
		visible:     true,
		initialized: true,
	}
}

func NewComponent(tpl []byte) *Component {
	return &Component{
		template: tpl,
	}
}

// Parent return the parent component
func (c *Component) Parent() Elm {
	return c.parent
}

// Children return all child components
func (c *Component) Children() []Elm {
	return c.children
}

func (c *Component) SetParent(el Elm) (err error) {
	if c.initialized {
		return fmt.Errorf("cannot set parent on initialized component")
	}
	c.parent = el
	return
}

// SetChildren initialize the child components
func (c *Component) SetChildren(elms ...Elm) (err error) {
	for _, el := range elms {
		if err = el.SetParent(c); err != nil {
			return
		}
	}
	c.children = elms
	return
}

// Focused change the focus state of component to true
func (c *Component) Focused() error {
	c.focused = true
	return nil
}

// Unfocus change the focus state of component to false
func (c *Component) Unfocus() error {
	for _, v := range c.children {
		if f, ok := v.(Focusable); ok {
			_ = f.Unfocus()
		}
	}
	c.focused = false
	return nil
}

// IsFocused return the current focus state of component
func (c *Component) IsFocused() bool {
	return c.focused
}

// Show change the visibility state of component to true
func (c *Component) Show() error {
	c.visible = true
	return nil
}

// Hide change the visibility state of component to false
func (c *Component) Hide() error {
	c.visible = false
	return nil
}

// IsVisible return the visible state of component
func (c *Component) IsVisible() bool {
	return c.visible
}

// IsRoot determine for root component
func (c *Component) IsRoot() bool {
	return c.parent == nil && c.initialized
}

// Init inherit from tea
func (c *Component) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	for _, el := range c.children {
		cmds = append(cmds, el.Init())
	}

	return tea.Batch(cmds...)
}

// Init inherit from tea
func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)
	for _, el := range c.children {
		_, cmd := el.Update(msg)
		cmds = append(cmds, cmd)
	}
	return c, tea.Batch(cmds...)
}

// Init inherit from tea
func (c *Component) View() string {
	return ""
}

func (c *Component) SubView() string {
	var view string
	for _, el := range c.children {
		view += el.View()
	}
	return view
}
