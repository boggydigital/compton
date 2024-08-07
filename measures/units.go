package measures

type Unit int

const (
	Small Unit = iota
	Normal
	Large
)

var unitCustomProperties = map[Unit]string{
	Small:  "--small",
	Normal: "--normal",
	Large:  "--large",
}

func (u Unit) String() string {
	return unitCustomProperties[u]
}
