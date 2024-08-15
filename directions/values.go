package directions

type Direction int

const (
	Column Direction = iota
	Row
)

var directionStrings = map[Direction]string{
	Column: "column",
	Row:    "row",
}

func (d Direction) String() string {
	return directionStrings[d]
}
