package vui

type HtmlElm struct {
	*Elm
}

var (
	_ Elementary = (*HtmlElm)(nil)
)

func (e *HtmlElm) New(opts ...interface{}) Elementary {
	return &HtmlElm{
		Elm: NewElm(opts...),
	}
}
