package measures

type Unit int

const (
	Small Unit = iota
	Normal
	Large
)

var unitStrings = map[Unit]string{
	Small:  "small",
	Normal: "normal",
	Large:  "large",
}

func (u Unit) String() string {
	return unitStrings[u]
}
