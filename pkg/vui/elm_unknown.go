package vui

type UnknownElm struct {
	*Elm
}

var (
	_ Elementary = (*HtmlElm)(nil)
)

func (e *UnknownElm) New(opts ...interface{}) Elementary {
	return &UnknownElm{
		Elm: NewElm(opts...),
	}
}
