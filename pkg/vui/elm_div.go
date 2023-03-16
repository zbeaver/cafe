package vui

type DivElm struct {
	*Elm
}

var (
	_ Elementary = (*DivElm)(nil)
)

func (e *DivElm) New(opts ...interface{}) Elementary {
	return &DivElm{
		Elm: NewElm(opts...),
	}
}
