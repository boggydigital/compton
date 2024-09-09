package section

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "section"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/section.html"
	markupCSection []byte
	//go:embed "style/section.css"
	styleCSection []byte
)

type SectionElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (se *SectionElement) WriteStyles(w io.Writer) error {
	if se.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleCSection).WriteContent(w); err != nil {
			return err
		}
	}
	return nil
}

func Section(r compton.Registrar) compton.Element {
	return &SectionElement{
		BaseElement: compton.BaseElement{
			Markup:  markupCSection,
			TagName: compton_atoms.CSection,
		},
		r: r,
	}
}
