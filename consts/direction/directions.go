package direction

type Direction int

const (
	Unset Direction = iota
	Column
	ColumnReverse
	Row
	RowReverse
)

var directionStrings = map[Direction]string{
	Column:        "column",
	ColumnReverse: "column-reverse",
	Row:           "row",
	RowReverse:    "row-reverse",
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
	return Unset
}
