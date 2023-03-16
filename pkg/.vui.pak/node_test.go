package vui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_Simple(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, 1)
}
