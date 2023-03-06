package vui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/ryboe/q"
)

type Elm interface {
	tea.Model
	SetParent(Elm) error
	SetChildren(...Elm) error
	AppendChild(...Elm) error
}

type Focusable interface {
	Focused() error
	Unfocus() error
	IsFocused() bool
}

type Switcher interface {
	Show() error
	Hide() error
}

type component struct {
	uid string
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
}

var (
	_ Focusable = (*component)(nil)
	_ Switcher  = (*component)(nil)
)

// NewRootcomponent return a component root
func NewRootComponent() *component {
	id, _ := uuid.NewUUID()
	return &component{
		uid:     id.String(),
		parent:  nil,
		focused: true,
		visible: true,
	}
}

func NewComponent() *component {
	id, _ := uuid.NewUUID()
	q.Q(id)
	return &component{
		uid: id.String(),
	}
}

// Parent return the parent component
func (c *component) Parent() Elm {
	return c.parent
}

// Children return all child components
func (c *component) Children() []Elm {
	return c.children
}

func (c *component) SetParent(el Elm) (err error) {
	if c.initialized {
		return fmt.Errorf("cannot set parent on initialized component")
	}
	c.parent = el
	return
}

// SetChildren initialize the child components
func (c *component) SetChildren(elms ...Elm) (err error) {
	for _, el := range elms {
		if err = el.SetParent(c); err != nil {
			return
		}
	}
	c.children = elms
	return
}

func (c *component) AppendChild(elms ...Elm) (err error) {
	for _, el := range elms {
		if err = el.SetParent(c); err != nil {
			return
		}
	}
	c.children = append(c.children, elms...)
	return
}

// Focused change the focus state of component to true
func (c *component) Focused() error {
	c.focused = true
	return nil
}

// Unfocus change the focus state of component to false
func (c *component) Unfocus() error {
	for _, v := range c.children {
		if f, ok := v.(Focusable); ok {
			_ = f.Unfocus()
		}
	}
	c.focused = false
	return nil
}

// IsFocused return the current focus state of component
func (c *component) IsFocused() bool {
	return c.focused
}

// Show change the visibility state of component to true
func (c *component) Show() error {
	c.visible = true
	return nil
}

// Hide change the visibility state of component to false
func (c *component) Hide() error {
	c.visible = false
	return nil
}

// IsVisible return the visible state of component
func (c *component) IsVisible() bool {
	return c.visible
}

// IsRoot determine for root component
func (c *component) IsRoot() bool {
	return c.parent == nil && c.initialized
}

// Init inherit from tea
func (c *component) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	for _, el := range c.children {
		cmds = append(cmds, el.Init())
	}

	return tea.Batch(cmds...)
}

// Init inherit from tea
func (c *component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (c *component) View() string {
	return ""
}

// @TODO remove it
func (c *component) SubView() string {
	var view string
	for _, el := range c.children {
		view += el.View()
	}
	return view
}
