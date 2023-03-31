package vui

type OptGroupElm struct {
	*Elm
}

var (
	_ Elementary = (*OptGroupElm)(nil)
)

func (e *OptGroupElm) New(opts ...interface{}) Elementary {
	return &OptGroupElm{
		Elm: NewElm(opts...),
	}
}
