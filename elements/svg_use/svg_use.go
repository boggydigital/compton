package svg_use

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"io"
)

type Symbol int

const (
	None Symbol = iota
	Windows
	MacOS
	Linux
	Plus
	Star
	Sparkle
	Stack
	Search
	Circle
)

var symbolStrings = map[Symbol]string{
	Windows: "windows",
	MacOS:   "macos",
	Linux:   "linux",
	Plus:    "plus",
	Star:    "star",
	Sparkle: "sparkle",
	Stack:   "stack",
	Search:  "search",
	Circle:  "circle",
}

const registrationName = "symbols"

var (
	//go:embed "markup/atlas.html"
	markupAtlas []byte
	//go:embed "markup/svg-use.html"
	markupSvgUse []byte
)

type SvgUseElement struct {
	compton.BaseElement
	r compton.Registrar
	s Symbol
}

func (sue *SvgUseElement) WriteRequirements(w io.Writer) error {
	if sue.r.RequiresRegistration(registrationName) {
		if _, err := w.Write(markupAtlas); err != nil {
			return err
		}
	}
	return nil
}

func (sue *SvgUseElement) WriteContent(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupSvgUse), w, sue.fragmentWriter)
}

func (sue *SvgUseElement) fragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Symbol":
		if _, err := io.WriteString(w, symbolStrings[sue.s]); err != nil {
			return err
		}
	case compton.AttributesToken:
		if err := sue.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func SvgUse(r compton.Registrar, s Symbol) compton.Element {
	sue := &SvgUseElement{
		r: r,
		s: s,
	}
	sue.AddClass(symbolStrings[s])
	return sue
}
