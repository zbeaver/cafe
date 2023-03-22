package vui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCss_Parse(t *testing.T) {
	as := assert.New(t)
	css := NewCSSStyleDecl(`color:red;background:green;border: 1px dashed blue`)
	as.Equal("red", css.GetPropertyValue("color"))
	as.Equal("green", css.GetPropertyValue("background"))
	as.Equal("1px dashed blue", css.GetPropertyValue("border"))
}
