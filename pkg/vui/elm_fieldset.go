package vui

type FieldSetElm struct {
	*Elm
}

var (
	_ Elementary = (*FieldSetElm)(nil)
)

func (e *FieldSetElm) New(opts ...interface{}) Elementary {
	return &FieldSetElm{
		Elm: NewElm(opts...),
	}
}
