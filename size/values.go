package size

type Size int

const (
	Normal Size = iota

	Small
	XSmall
	XXSmall
	XXXSmall

	Large
	XLarge
	XXLarge
	XXXLarge
)

var unitStrings = map[Size]string{
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

func (u Size) String() string {
	return unitStrings[u]
}
