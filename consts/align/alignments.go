package align

type Align int

const (
	Unset Align = iota
	Start
	Center
	End
	SpaceBetween
	SpaceAround
	SpaceEvenly
)

var alignStrings = map[Align]string{
	Start:        "start",
	Center:       "center",
	End:          "end",
	SpaceBetween: "space-between",
	SpaceAround:  "space-around",
	SpaceEvenly:  "space-evenly",
}

func (a Align) String() string {
	return alignStrings[a]
}

func Parse(s string) Align {
	for a, str := range alignStrings {
		if str == s {
			return a
		}
	}
	return Unset
}
