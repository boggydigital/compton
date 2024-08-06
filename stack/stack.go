package stack

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/custom_elements"
	"golang.org/x/net/html/atom"
	"io"
)

type Gap int

const (
	Small Gap = iota
	Normal
	Large
)

const (
	// Atom for stack is the first value created,
	// using max value and leaving 255 more possible atoms
	Atom atom.Atom = 0xffffff00
)

var gapCustomProperties = map[Gap]string{
	Small:  "--small",
	Normal: "--normal",
	Large:  "--large",
}

const (
	elementName = "c-stack"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/c-stack.html"
	markupStack []byte
)

type Stack struct {
	compton.BaseElement
	gap Gap
	wcr compton.Registrar
}

func (s *Stack) Register(w io.Writer) error {
	if s.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, s.templateFragmentWriter); err != nil {
			return err
		}
	}
	return nil
}

func (s *Stack) Write(w io.Writer) error {
	return s.BaseElement.Write(w)
}

func (s *Stack) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Spacing":
		if _, err := io.WriteString(w, gapCustomProperties[s.gap]); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func New(wcr compton.Registrar, gap Gap) compton.Element {
	return &Stack{
		BaseElement: compton.BaseElement{
			Markup:  markupStack,
			TagName: Atom,
		},
		wcr: wcr,
		gap: gap,
	}
}
