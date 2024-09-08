package color

import (
	"iter"
	"maps"
)

type Color int

const (
	Black Color = iota
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
	DeepPurple: "deep-purple",
	Indigo:     "indigo",
	Blue:       "blue",
	LightBlue:  "light-blue",
	Cyan:       "cyan",
	Teal:       "teal",
	Green:      "green",
	LightGreen: "light-green",
	Lime:       "lime",
	Yellow:     "yellow",
	Amber:      "amber",
	Orange:     "orange",
	DeepOrange: "deep-orange",
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
