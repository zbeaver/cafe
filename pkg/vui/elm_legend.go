package vui

type LegendElm struct {
	*Elm
}

var (
	_ Elementary = (*LegendElm)(nil)
)

func (e *LegendElm) New(opts ...interface{}) Elementary {
	return &LegendElm{
		Elm: NewElm(opts...),
	}
}
