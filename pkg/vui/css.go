package vui

import (
	"bytes"

	parse "github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type CSSStyleSheet struct{}
type CSSRuleList []CSSRule
type CSSRule struct{}
type CSSStyleDecl struct {
	Text string
	css  map[string]string
}

func NewCSSStyleDecl(text string) CSSStyleDecl {
	p := css.NewParser(parse.NewInput(bytes.NewBufferString(text)), true)
	out := map[string]string{}
	for {
		gt, _, data := p.Next()
		if gt == css.ErrorGrammar {
			break
		}
		if gt == css.AtRuleGrammar ||
			gt == css.BeginAtRuleGrammar ||
			gt == css.BeginRulesetGrammar ||
			gt == css.DeclarationGrammar {
			for _, val := range p.Values() {
				out[string(data)] += string(val.Data)
			}
		}
	}

	return CSSStyleDecl{
		Text: text,
		css:  out,
	}
}

func (s CSSStyleDecl) GetPropertyValue(name string) string {
	val, ok := s.css[name]
	if ok {
		return val
	}
	return ""
}
