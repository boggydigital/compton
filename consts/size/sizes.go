package size

type Size int

const (
	Unknown Size = iota

	Normal

	Small
	XSmall
	XXSmall
	XXXSmall

	Large
	XLarge
	XXLarge
	XXXLarge
)

var sizeStrings = map[Size]string{
	Normal:   "n",
	Small:    "s",
	XSmall:   "xs",
	XXSmall:  "xxs",
	XXXSmall: "xxxs",
	Large:    "l",
	XLarge:   "xl",
	XXLarge:  "xxl",
	XXXLarge: "xxxl",
}

func (u Size) String() string {
	return sizeStrings[u]
}

func (u Size) CssValue() string {
	return "var(--s-" + u.String() + ")"
}

func Parse(s string) Size {
	for sz, str := range sizeStrings {
		if s == str {
			return sz
		}
	}
	return Unknown
}
