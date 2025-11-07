package size

type Size int

const (
	Unset Size = iota
	Initial
	Zero
	Normal
	Small
	XSmall
	XXSmall
	XXXSmall
	Large
	XLarge
	XXLarge
	XXXLarge

	ColumnWidth
	MaxWidth
	FullWidth
	FitContent
)

var sizeStrings = map[Size]string{
	Unset:       "unset",
	Initial:     "initial",
	Zero:        "zero",
	Normal:      "n",
	Small:       "s",
	XSmall:      "xs",
	XXSmall:     "xxs",
	XXXSmall:    "xxxs",
	Large:       "l",
	XLarge:      "xl",
	XXLarge:     "xxl",
	XXXLarge:    "xxxl",
	ColumnWidth: "cw",
	MaxWidth:    "maxw",
	FullWidth:   "fw",
	FitContent:  "fc",
}

func (u Size) String() string {
	return sizeStrings[u]
}

func (u Size) SizeCssValue() string {
	return "var(--s-" + u.String() + ")"
}

func (u Size) FontSizeCssValue() string {
	return "var(--fs-" + u.String() + ")"
}

func Parse(s string) Size {
	for sz, str := range sizeStrings {
		if s == str {
			return sz
		}
	}
	return Unset
}
