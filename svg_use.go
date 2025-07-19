package compton

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"io"
)

type Symbol int

const (
	NoSymbol Symbol = iota
	Windows
	MacOS
	Linux
	Sparkle
	Stack
	Search
	Circle
	TwoDownwardChevrons
	DownwardNestedChevrons
	UpwardNestedChevrons
	CompactDisk
	RisingSun
	NewsBroadcast
	ShoppingLabel
	Bookmark
	Percent
)

var symbolStrings = map[Symbol]string{
	Windows:                "windows",
	MacOS:                  "macos",
	Linux:                  "linux",
	Sparkle:                "sparkle",
	Stack:                  "stack",
	Search:                 "search",
	Circle:                 "circle",
	TwoDownwardChevrons:    "two-downward-chevrons",
	DownwardNestedChevrons: "downward-nested-chevrons",
	UpwardNestedChevrons:   "upward-nested-chevrons",
	CompactDisk:            "compact-disk",
	RisingSun:              "rising-sun",
	NewsBroadcast:          "news-broadcast",
	ShoppingLabel:          "shopping-label",
	Bookmark:               "bookmark",
	Percent:                "percent",
}

var (
	//go:embed "markup/atlas.html"
	markupAtlas string
)

type SvgUseElement struct {
	*BaseElement
	s Symbol
}

func (sue *SvgUseElement) Write(w io.Writer) error {
	bts, err := sue.BaseElement.MarkupProvider.GetMarkup()
	if err != nil {
		return err
	}
	return WriteContents(bytes.NewReader(bts), w, sue.WriteFragment)
}

func (sue *SvgUseElement) WriteFragment(t string, w io.Writer) error {
	switch t {
	case ".Symbol":
		if _, err := io.WriteString(w, symbolStrings[sue.s]); err != nil {
			return err
		}
	case AttributesToken:
		if err := sue.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func (sue *SvgUseElement) ForegroundColor(c color.Color) *SvgUseElement {
	sue.AddClass(class.ForegroundColor(c))
	return sue
}

func SvgUse(r Registrar, s Symbol) *SvgUseElement {
	sue := &SvgUseElement{
		BaseElement: NewElement(atomsEmbedMarkup(compton_atoms.SvgUse, DefaultMarkup)),
		s:           s,
	}

	sue.AddClass(symbolStrings[s])

	r.RegisterStyles(DefaultStyle, compton_atoms.StyleName(compton_atoms.SvgUse))
	r.RegisterRequirements(compton_atoms.MarkupName(compton_atoms.SvgUse), Text(markupAtlas))

	return sue
}
