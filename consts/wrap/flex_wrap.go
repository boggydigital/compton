package wrap

type FlexWrap int

const (
	NoWrap FlexWrap = iota
	Wrap
	WrapReverse
)

var flexWrapStrings = map[FlexWrap]string{
	NoWrap:      "nowrap",
	Wrap:        "wrap",
	WrapReverse: "wrap-reverse",
}

func (fw FlexWrap) String() string {
	return flexWrapStrings[fw]
}

func Parse(s string) FlexWrap {
	for fw, str := range flexWrapStrings {
		if s == str {
			return fw
		}
	}
	return Wrap
}
