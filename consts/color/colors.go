package color

import (
	"maps"
	"slices"
)

type Color int

const (
	Unset Color = iota
	Black
	White
	Red
	Orange
	Yellow
	Green
	Mint
	Teal
	Cyan
	Blue
	Indigo
	Purple
	Pink
	Brown
	Background
	Foreground
	DimmedForeground
	Gray
	Highlight
	Transparent
	Inherit
	Rep
	RepForeground
	RepBackground
	RepGray
)

var colorStrings = map[Color]string{
	Black:            "black",
	White:            "white",
	Red:              "red",
	Orange:           "orange",
	Yellow:           "yellow",
	Green:            "green",
	Mint:             "mint",
	Teal:             "teal",
	Cyan:             "cyan",
	Blue:             "blue",
	Indigo:           "indigo",
	Purple:           "purple",
	Pink:             "pink",
	Brown:            "brown",
	Gray:             "gray",
	Background:       "background",
	Foreground:       "foreground",
	DimmedForeground: "dimmedforeground",
	Highlight:        "highlight",
	Transparent:      "transparent",
	Inherit:          "inherit",
	Rep:              "rep",
	RepForeground:    "repforeground",
	RepBackground:    "repbackground",
	RepGray:          "repgray",
}

func (c Color) String() string {
	return colorStrings[c]
}

func (c Color) CssValue() string {
	return "var(--c-" + c.String() + ")"
}

func Parse(s string) Color {
	for c, str := range colorStrings {
		if s == str {
			return c
		}
	}
	return Unset
}

func All() []Color {
	return slices.Collect(maps.Keys(colorStrings))
}
