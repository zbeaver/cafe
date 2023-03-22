package render

import (
	"context"
	"testing"

	"github.com/ryboe/q"
	"github.com/stretchr/testify/assert"
	"github.com/zbeaver/cafe/pkg/decoder"
	"github.com/zbeaver/cafe/pkg/vui"
)

type testcase struct {
	tpl      []byte
	doc      vui.Documentary
	expected func(*assert.Assertions, string)
}

func TestRender_Simple(t *testing.T) {
	as := assert.New(t)
	tc := []testcase{
		{
			[]byte(`
				<html>
					<body style="background-color: #352456">
					<div name="div1">
						xxx
						<div>world</div>
					</div>
					<div name="div2">
						hehe
						<div> coucu</div>
					</div>
					</body>
				</html>
		 `),
			vui.NewDocument(),
			func(as *assert.Assertions, result string) {
				q.Q(result)
			},
		},
	}

	for _, c := range tc {
		de := decoder.NewDecoder()
		_, err := de.Decode(c.doc, c.tpl)
		if err != nil {
			panic(err)
		}
		engine := NewEngine(context.Background(), c.doc)
		c.expected(as, engine.Render())
	}
}
