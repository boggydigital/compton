package input_types

type Type int

const (
	Search Type = iota
	Text
	Submit
	Button
	Checkbox
	Hidden
)

var inputTypeStrings = map[Type]string{
	Search:   "search",
	Text:     "text",
	Submit:   "submit",
	Button:   "button",
	Checkbox: "checkbox",
	Hidden:   "hidden",
}

func (it Type) String() string {
	return inputTypeStrings[it]
}
