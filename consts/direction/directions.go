package direction

type Direction int

const (
	Unknown Direction = iota
	Column
	Row
)

var directionStrings = map[Direction]string{
	Column: "column",
	Row:    "row",
}

func (d Direction) String() string {
	return directionStrings[d]
}

func Parse(s string) Direction {
	for d, str := range directionStrings {
		if str == s {
			return d
		}
	}
	return Unknown
}
