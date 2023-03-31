package vui

type SpanElm struct {
	*Elm
}

var (
	_ Elementary = (*SpanElm)(nil)
)

func (e *SpanElm) New(opts ...interface{}) Elementary {
	return &SpanElm{
		Elm: NewElm(opts...),
	}
}
