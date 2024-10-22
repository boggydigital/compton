package compton

import (
	"bytes"
	_ "embed"
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
	BaseElement
	//r compton.Registrar
	s Symbol
}

//func (sue *SvgUseElement) WriteRequirements(w io.Writer) error {
//	if sue.r.RequiresRegistration(registrationName) {
//		if _, err := w.Write(markupAtlas); err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (sue *SvgUseElement) Write(w io.Writer) error {
	return WriteContents(bytes.NewReader(markupSvgUse), w, sue.fragmentWriter)
}

func (sue *SvgUseElement) fragmentWriter(t string, w io.Writer) error {
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

func SvgUse(r Registrar, s Symbol) Element {
	sue := &SvgUseElement{
		//r: r,
		s: s,
	}
	sue.AddClass(symbolStrings[s])

	r.RegisterRequirement(registrationName, TextBytes(markupAtlas))

	return sue
}
