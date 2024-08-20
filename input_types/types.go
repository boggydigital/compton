package input_types

type Type int

const (
	Search Type = iota
	Text
)

var inputTypeStrings = map[Type]string{
	Search: "search",
	Text:   "text",
}

func (it Type) String() string {
	return inputTypeStrings[it]
}
