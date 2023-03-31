package vui

type InputElm struct {
	*Elm
}

var (
	_ Elementary = (*InputElm)(nil)
)

func (e *InputElm) New(opts ...interface{}) Elementary {
	return &InputElm{
		Elm: NewElm(opts...),
	}
}
