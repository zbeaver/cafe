package vui

type HeadingElm struct {
	*Elm
	align string
}

var (
	_ Elementary = (*HeadingElm)(nil)
)

func (e *HeadingElm) New(opts ...interface{}) Elementary {
	return &HeadingElm{
		Elm: NewElm(opts...),
	}
}
