package compton

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

type Symbol int

const (
	NoSymbol Symbol = iota
	Windows
	PlaceOfInterest
	SteamDeck
	Sparkle
	Search
	CirclePositive
	TriangleNeutral
	CrossNegative
	UpwardChevron
	RightwardChevron
	DownwardChevron
	LeftwardChevron
	ThreeDownwardChevrons
	ThreeUpwardChevrons
	ThreeHorizontalLines
	CircleCompactDisk
	CircleDashed
	CircleClockArrows
	CircleDownwardArrow
	RisingSun
	NewsBroadcast
	ShoppingLabel
	Bookmark
	Percent
	Heart
	Shield
	PuzzlePiece
	TwoTitleValues
	Gemstone
	TwoStackedItems
	ItemPlus
	Trophy
	HexagonDiagonalLines
	HexagonClockArrows
	HexagonDownwardArrow
	HexagonSparkling
	VideoThumbnail
	ImageThumbnail
	ExternalLink
)

var symbolStrings = map[Symbol]string{
	Windows:               "windows",
	PlaceOfInterest:       "place-of-interest",
	SteamDeck:             "steam-deck",
	Sparkle:               "sparkle",
	Search:                "search",
	CirclePositive:        "circle-positive",
	TriangleNeutral:       "triangle-neutral",
	CrossNegative:         "cross-negative",
	UpwardChevron:         "upward-chevron",
	RightwardChevron:      "rightward-chevron",
	DownwardChevron:       "downward-chevron",
	LeftwardChevron:       "leftward-chevron",
	ThreeDownwardChevrons: "three-downward-chevrons",
	ThreeUpwardChevrons:   "three-upward-chevrons",
	ThreeHorizontalLines:  "three-horizontal-lines",
	CircleDashed:          "circle-dashed",
	CircleClockArrows:     "circle-clock-arrows",
	CircleDownwardArrow:   "circle-downward-arrow",
	CircleCompactDisk:     "compact-disk",
	RisingSun:             "rising-sun",
	NewsBroadcast:         "news-broadcast",
	ShoppingLabel:         "shopping-label",
	Bookmark:              "bookmark",
	Percent:               "percent",
	Heart:                 "heart",
	Shield:                "shield",
	PuzzlePiece:           "puzzle-piece",
	TwoTitleValues:        "two-title-values",
	Gemstone:              "gemstone",
	TwoStackedItems:       "two-stacked-items",
	ItemPlus:              "item-plus",
	Trophy:                "trophy",
	HexagonDiagonalLines:  "hexagon-diagonal-lines",
	HexagonClockArrows:    "hexagon-clock-arrows",
	HexagonDownwardArrow:  "hexagon-downward-arrow",
	HexagonSparkling:      "hexagon-sparkling",
	VideoThumbnail:        "video-thumbnail",
	ImageThumbnail:        "image-thumbnail",
	ExternalLink:          "external-link",
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
