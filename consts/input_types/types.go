package input_types

type Type int

const (
	Search Type = iota
	Text
	Submit
	Button
	Checkbox
	Hidden
	Password
)

var inputTypeStrings = map[Type]string{
	Search:   "search",
	Text:     "text",
	Submit:   "submit",
	Button:   "button",
	Checkbox: "checkbox",
	Hidden:   "hidden",
	Password: "password",
}

func (it Type) String() string {
	return inputTypeStrings[it]
}
