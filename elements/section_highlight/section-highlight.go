package section_highlight

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"io"
)

const (
	elementName = "section-highlight"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/section-highlight.html"
	markupSectionHighlight []byte
)

type SectionHighlightElement struct {
	compton.BaseElement
	wcr compton.Registrar
}

func (sh *SectionHighlightElement) WriteRequirements(w io.Writer) error {
	if sh.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return sh.BaseElement.WriteRequirements(w)
}

func SectionHighlight(wcr compton.Registrar) compton.Element {
	return &SectionHighlightElement{
		BaseElement: compton.BaseElement{
			Markup:  markupSectionHighlight,
			TagName: compton_atoms.SectionHighlight,
		},
		wcr: wcr,
	}
}
