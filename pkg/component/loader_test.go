package vui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simplest case
// The expectation of this test case:
// Scan template content and initialize all child components found

// Test plain component
type T_Plain struct {
	Elm
}

func (c *T_Plain) Template() []byte {
	return []byte(`
	<template> hard code </template>
	`)
}

func (c *T_Plain) Components() Registry {
	return Registry{}
}

func TestLoader_Plain(t *testing.T) {
	s := assert.New(t)
	a := &T_Plain{
		Elm: NewComponent(),
	}
	loader := NewLoader(a)
	if err := loader.Load(); err != nil {
		panic(err)
	}
	s.Equal(a.View(), "hard code")
}

// Test elm - templateless

// Test slot case
type T_AppZ struct {
	Elm
}

func (c *T_AppZ) Template() []byte {
	return []byte(`
	<template>
		<text>
			<text>
				<text>
					<text>hel<input>asdf</input>lo</text>
				</text>
			</text>
		</text>
		<text>
			<text>there</text>
		</text>
	</template>
	`)
}

func (c *T_AppZ) Components() Registry {
	return Registry{
		"text": func() Loadable { return &T_Text{Elm: NewComponent()} },
	}
}

func (c *T_Text) Template() []byte {
	return []byte("<template>hello</template>")
}

func (c *T_Text) Components() Registry {
	return Registry{}
}

type T_Text struct {
	Elm
}

func TestLoader_Nested(t *testing.T) {
	s := assert.New(t)
	root := &T_AppZ{
		Elm: NewComponent(),
	}
	loader := NewLoader(root)
	if err := loader.Load(); err != nil {
		panic(err)
	}
	s.Equal(root.View(), "hello world")
}
