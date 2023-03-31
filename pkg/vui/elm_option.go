package vui

type OptionElm struct {
	*Elm
}

var (
	_ Elementary = (*OptionElm)(nil)
)

func (e *OptionElm) New(opts ...interface{}) Elementary {
	return &OptionElm{
		Elm: NewElm(opts...),
	}
}
