package vui

type HrElm struct {
	*Elm
}

var (
	_ Elementary = (*HrElm)(nil)
)

func (e *HrElm) New(opts ...interface{}) Elementary {
	return &HrElm{
		Elm: NewElm(opts...),
	}
}
