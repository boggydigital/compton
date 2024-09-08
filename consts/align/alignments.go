package align

import (
	"iter"
	"maps"
	"strings"
)

type Align int

const (
	Unknown Align = iota
	Start
	Center
	End
)

var alignStrings = map[Align]string{
	Start:  "start",
	Center: "center",
	End:    "end",
}

func (a Align) String() string {
	return alignStrings[a]
}

func Parse(s string) Align {
	for a, str := range alignStrings {
		if strings.HasPrefix(str, s) {
			return a
		}
	}
	return Unknown
}

func AllAligns() iter.Seq[Align] {
	return maps.Keys(alignStrings)
}
