package vui

var (
	_ INode = (*Text)(nil)
)

type Texter interface {
	INode
}

type Text struct {
	*Node
}

func (t *Text) New(val string) *Text {
	return &Text{
		Node: NewNode(
			WithNodeName("text", TextNode),
			WithNodeValue(val),
		),
	}
}
