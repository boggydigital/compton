package input_types

type Type int

const (
	Search Type = iota
	Text
	Submit
)

var inputTypeStrings = map[Type]string{
	Search: "search",
	Text:   "text",
	Submit: "submit",
}

func (it Type) String() string {
	return inputTypeStrings[it]
}
