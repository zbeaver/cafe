package vui

type ParagraphElm struct {
	*Elm
}

var (
	_ Elementary = (*ParagraphElm)(nil)
)

func (e *ParagraphElm) New(opts ...interface{}) Elementary {
	return &ParagraphElm{
		Elm: NewElm(opts...),
	}
}
