package vui

type FormElm struct {
	*Elm
}

var (
	_ Elementary = (*FormElm)(nil)
)

func (e *FormElm) New(opts ...interface{}) Elementary {
	return &FormElm{
		Elm: NewElm(opts...),
	}
}
