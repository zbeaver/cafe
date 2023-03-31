package vui

type AnchorElm struct {
	*Elm
}

var (
	_ Elementary = (*AnchorElm)(nil)
)

func (e *AnchorElm) New(opts ...interface{}) Elementary {
	return &AnchorElm{
		Elm: NewElm(opts...),
	}
}
