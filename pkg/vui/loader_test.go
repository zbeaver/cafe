package vui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simplest case
type T_AppZ struct {
	Elm
}

func (c *T_AppZ) Template() []byte {
	return []byte(`
	<template>
		<text>
			<text>
				<text>
					<text>hello</text>
				</text>
			</text>
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
	return []byte("")
}

func (c *T_Text) Components() Registry {
	return Registry{}
}

type T_Text struct {
	Elm
}

func TestLoader_Simplest(t *testing.T) {
	s := assert.New(t)
	root := &T_AppZ{
		Elm: NewComponent(),
	}
	loader := NewLoader(root)
	if err := loader.Load(); err != nil {
		fmt.Errorf("oops unmarshall failed")
	}
	s.Equal("hello", "hello")
}
