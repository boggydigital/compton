package weight

type Weight int

const (
	Unknown Weight = iota

	Normal

	Bolder
	Lighter
)

var weightStrings = map[Weight]string{
	Normal:  "n",
	Bolder:  "b",
	Lighter: "l",
}

func (w Weight) String() string {
	return weightStrings[w]
}

func (w Weight) CssValue() string {
	return "var(--fw-" + w.String() + ")"
}

func Parse(w string) Weight {
	for wt, str := range weightStrings {
		if w == str {
			return wt
		}
	}
	return Unknown
}
