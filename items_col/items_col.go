package items_col

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/shared"
	"io"
)

const (
	elementName = "items-col"
	gapAttr     = "data-gap"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/items-col.html"
	markupItemsCol []byte
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
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, s.templateFragmentWriter); err != nil {
			return err
		}
	}
	return s.BaseElement.Register(w)
}

func (s *Stack) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".HostGaps":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostGaps)); err != nil {
			return err
		}
	}
	return nil
}

func (s *Stack) SetGap(amount measures.Unit) *Stack {
	s.SetAttr(gapAttr, amount.String())
	return s
}

func New(wcr compton.Registrar) *Stack {
	return &Stack{
		BaseElement: compton.BaseElement{
			Markup:  markupItemsCol,
			TagName: compton_atoms.ItemsCol,
		},
		wcr: wcr,
	}
}
