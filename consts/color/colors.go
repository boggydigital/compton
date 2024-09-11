package color

import (
	_ "embed"
	"iter"
	"maps"
)

//go:embed "style/colors.css"
var StyleSheet []byte

type Color int

const (
	Unset Color = iota
	Black
	White
	Red
	Pink
	Purple
	DeepPurple
	Indigo
	Blue
	LightBlue
	Cyan
	Teal
	Green
	LightGreen
	Lime
	Yellow
	Amber
	Orange
	DeepOrange
	Brown
	Background
	Foreground
	Subtle
	Highlight
	Shadow
)

var colorStrings = map[Color]string{
	Black:      "black",
	White:      "white",
	Red:        "red",
	Pink:       "pink",
	Purple:     "purple",
	DeepPurple: "deeppurple",
	Indigo:     "indigo",
	Blue:       "blue",
	LightBlue:  "lightblue",
	Cyan:       "cyan",
	Teal:       "teal",
	Green:      "green",
	LightGreen: "lightgreen",
	Lime:       "lime",
	Yellow:     "yellow",
	Amber:      "amber",
	Orange:     "orange",
	DeepOrange: "deeporange",
	Brown:      "brown",
	Background: "background",
	Foreground: "foreground",
	Subtle:     "subtle",
	Highlight:  "highlight",
	Shadow:     "shadow",
}

func (c Color) String() string {
	return colorStrings[c]
}

func (c Color) CssValue() string {
	return "var(--c-" + c.String() + ")"
}

func AllColors() iter.Seq[Color] {
	return maps.Keys(colorStrings)
}

func Parse(s string) Color {
	for c, str := range colorStrings {
		if s == str {
			return c
		}
	}
	return Unset
}
