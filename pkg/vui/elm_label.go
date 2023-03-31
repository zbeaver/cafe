package vui

type LabelElm struct {
	*Elm
}

var (
	_ Elementary = (*LabelElm)(nil)
)

func (e *LabelElm) New(opts ...interface{}) Elementary {
	return &LabelElm{
		Elm: NewElm(opts...),
	}
}
