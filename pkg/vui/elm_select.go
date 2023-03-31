package vui

type SelectElm struct {
	*Elm
}

var (
	_ Elementary = (*SelectElm)(nil)
)

func (e *SelectElm) New(opts ...interface{}) Elementary {
	return &SelectElm{
		Elm: NewElm(opts...),
	}
}
