package vui

type BodyElm struct {
	*Elm
}

var (
	_ Elementary = (*BodyElm)(nil)
)

func (e *BodyElm) New(opts ...interface{}) Elementary {
	return &BodyElm{
		Elm: NewElm(opts...),
	}
}
