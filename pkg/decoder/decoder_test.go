package decoder

import (
	"testing"

	"github.com/ryboe/q"
	"github.com/stretchr/testify/assert"
	"github.com/zbeaver/cafe/pkg/vui"
)

type testcase struct {
	tpl      []byte
	doc      vui.Documentary
	expected func(*assert.Assertions, vui.Elementary)
}

func TestDecoder(t *testing.T) {
	as := assert.New(t)
	tc := []testcase{
		{
			[]byte(`
				<div name="div1">
					xxx
					<div>world</div>
				</div>
				<div name="div2">
				  hehe
				  <div> coucu</div>
				</div>
			`),
			vui.NewDocument(),
			func(as *assert.Assertions, elm vui.Elementary) {
				as.Equal("#document", elm.NodeName())
				// HTML
				html := elm.ChildNodes()
				as.Equal(2, len(html))
				as.Equal("HEAD", html[0].NodeName())
				as.Equal("BODY", html[1].NodeName())

				// BODY
				body := html[1].ChildNodes()
				q.Q(body)

				// CONTENT
				as.Equal(4, len(body))
				as.Equal("DIV", body[0].NodeName())
				as.Equal("DIV", body[2].NodeName())

				// NESTED CONTENT
				div1 := body[0].ChildNodes()
				as.Equal(3, len(div1))
				as.Equal("#text", div1[0].NodeName())
				as.Equal("DIV", div1[1].NodeName())
				as.Equal("#text", div1[2].NodeName())

				div2 := body[2].ChildNodes()
				as.Equal(3, len(div2))
				as.Equal("#text", div2[0].NodeName())
				as.Equal("DIV", div2[1].NodeName())
				as.Equal("#text", div2[2].NodeName())
			},
		},
	}

	for _, c := range tc {
		de := NewDecoder()
		elm, err := de.Decode(c.doc, c.tpl)
		if err != nil {
			panic(err)
		}
		c.expected(as, elm)
	}
}
