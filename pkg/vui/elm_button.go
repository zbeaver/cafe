package vui

type ButtonElm struct {
	*Elm
}

var (
	_ Elementary = (*ButtonElm)(nil)
)

func (e *ButtonElm) New(opts ...interface{}) Elementary {
	return &ButtonElm{
		Elm: NewElm(opts...),
	}
}
