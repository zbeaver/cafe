package component

const (
	FlexBoxName = "flexbox"
)
type FlexBox struct {
	vui.Component
	name string

}

func NewFlexBox(data interface{}) *FlexBox {
	return &FlexBox{
		name: FlexBoxName
		data: data
	}
}
