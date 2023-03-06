package component

import "github.com/zbeaver/cafe/pkg/vui"

type Div struct {
	vui.Elm
	data struct {
		inline bool
	}
}

var (
	_ vui.Template = (*Div)(nil)
)

func (c *Div) Template() []byte {
	return []byte(`
		<root></root>
	`)
}

func (c *Div) New() Elm {
	return &Div{
		Elm: NewComponent()
	}
}

func (c *Div) Data() map[string]interface{} {
	return map[string]interface{

	}
}

func (c *Div) Prop() map[string]interface{} {
	return map[string]interface{
		"inline": true,
	}
}
