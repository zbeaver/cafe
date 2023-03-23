package vui

type BrElm struct {
	*Elm
}

var (
	_ Elementary = (*BrElm)(nil)
)

func (e *BrElm) New(opts ...interface{}) Elementary {
	return &BrElm{
		Elm: NewElm(opts...),
	}
}
