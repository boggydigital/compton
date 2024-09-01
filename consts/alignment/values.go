package alignment

type Position int

const (
	Start Position = iota
	Center
	End
)

var positionStrings = map[Position]string{
	Start:  "start",
	Center: "center",
	End:    "end",
}

func (p Position) String() string {
	return positionStrings[p]
}
