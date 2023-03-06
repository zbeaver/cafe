package vui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponent_ParentSetted(t *testing.T) {
	s := assert.New(t)
	root := NewRootComponent()
	c1 := &Component{}
	c2 := &Component{}
	root.SetChildren(c1, c2)
	s.Equal(root.IsRoot(), true)
	s.Equal(root, c1.Parent())
	s.Equal(root, c2.Parent())
	s.Equal(c1.initialized, false)
	s.Equal(c2.initialized, false)
}

func TestComponent_Unfocus(t *testing.T) {
	s := assert.New(t)
	root := NewRootComponent()
	l1 := &Component{}
	l2 := &Component{}
	l3 := &Component{}
	_ = l1.Focused()
	_ = l2.Focused()
	_ = l3.Focused()

	root.SetChildren(l1)
	l1.SetChildren(l2)
	l2.SetChildren(l3)
	_ = root.Unfocus()
	s.Equal(l1.IsFocused(), false)
	s.Equal(l2.IsFocused(), false)
	s.Equal(l3.IsFocused(), false)
}
