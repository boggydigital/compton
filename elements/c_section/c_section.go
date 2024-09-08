package c_section

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "c-section"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/c-section.html"
	markupCSection []byte
	//go:embed "style/c-section.css"
	styleCSection []byte
)

type CSectionElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (cs *CSectionElement) WriteStyles(w io.Writer) error {
	if cs.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleCSection).WriteContent(w); err != nil {
			return err
		}
	}
	return nil
}

func CSection(r compton.Registrar) compton.Element {
	return &CSectionElement{
		BaseElement: compton.BaseElement{
			Markup:  markupCSection,
			TagName: compton_atoms.CSection,
		},
		r: r,
	}
}
