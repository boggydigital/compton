package c_stack

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/measures"
	"golang.org/x/net/html/atom"
	"io"
)

const (
	// Atom for stack is the first value created,
	// using max value and leaving 255 more possible atoms
	Atom atom.Atom = 0xffffff00
)

const (
	elementName = "c-stack"
	rowGapAttr  = "data-row-gap"
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
}

func (s *Stack) Register(w io.Writer) error {
	if s.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return s.Parent.Register(w)
}

func (s *Stack) SetRowGap(amount measures.Unit) *Stack {
	s.SetAttr(rowGapAttr, amount.String())
	return s
}

func New(wcr compton.Registrar) *Stack {
	return &Stack{
		BaseElement: compton.BaseElement{
			Markup:  markupStack,
			TagName: Atom,
		},
		wcr: wcr,
	}
}
