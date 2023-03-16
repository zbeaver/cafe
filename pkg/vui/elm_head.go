package vui

type HeadElm struct {
	*Elm
}

var (
	_ Elementary = (*HeadElm)(nil)
)

func (e *HeadElm) New(opts ...interface{}) Elementary {
	return &HeadElm{
		Elm: NewElm(opts...),
	}
}
