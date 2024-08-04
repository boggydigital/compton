package stack

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
	"io"
)

type Gap int

const (
	Small Gap = iota
	Normal
	Large
)

var gapCustomProperties = map[Gap]string{
	Small:  "--small",
	Normal: "--normal",
	Large:  "--large",
}

const (
	elementName    = "c-stack"
	extendsElement = "HTMLElement"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/c-stack.html"
	markupStack []byte
)

type Stack struct {
	compton.BaseElement
	wcr compton.Registrar
	gap Gap
}

func (s *Stack) Write(w io.Writer) error {
	if err := s.wcr.RegisterCustomElement(elementName, extendsElement, compton.Closed, w); err != nil {
		return err
	}
	if err := s.wcr.RegisterMarkup(bytes.NewReader(markupTemplate), w, s.templateFragmentWriter); err != nil {
		return err
	}
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
			TagName: atom.Div,
		},
		wcr: wcr,
		gap: gap,
	}
}
