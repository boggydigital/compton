package measures

type Unit int

const (
	Normal Unit = iota

	Small
	XSmall
	XXSmall
	XXXSmall

	Large
	XLarge
	XXLarge
	XXXLarge
)

var unitStrings = map[Unit]string{
	Normal:   "normal",
	Small:    "small",
	XSmall:   "x-small",
	XXSmall:  "xx-small",
	XXXSmall: "xxx-small",
	Large:    "large",
	XLarge:   "x-large",
	XXLarge:  "xx-large",
	XXXLarge: "xxx-large",
}

func (u Unit) String() string {
	return unitStrings[u]
}
