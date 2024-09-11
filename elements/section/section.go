package section

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/consts/weight"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "section"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/section.html"
	markupSection []byte
	//go:embed "style/section.css"
	styleSection []byte
)

type SectionElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (se *SectionElement) WriteStyles(w io.Writer) error {
	if se.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleSection, styleRegistrationName).WriteContent(w); err != nil {
			return err
		}
	}
	return nil
}

func (se *SectionElement) BackgroundColor(c color.Color) *SectionElement {
	se.AddClass(class.BackgroundColor(c))
	return se
}

func (se *SectionElement) ForegroundColor(c color.Color) *SectionElement {
	se.AddClass(class.ForegroundColor(c))
	return se
}

func (se *SectionElement) FontSize(s size.Size) *SectionElement {
	se.AddClass(class.FontSize(s))
	return se
}

func (se *SectionElement) FontWeight(w weight.Weight) *SectionElement {
	se.AddClass(class.FontWeight(w))
	return se
}

func (se *SectionElement) RowGap(s size.Size) *SectionElement {
	se.AddClass(class.RowGap(s))
	return se
}

func (se *SectionElement) ColumnGap(s size.Size) *SectionElement {
	se.AddClass(class.ColumnGap(s))
	return se
}

func (se *SectionElement) Gap(s size.Size) *SectionElement {
	se.ColumnGap(s)
	se.RowGap(s)
	return se
}

func Section(r compton.Registrar) *SectionElement {
	return &SectionElement{
		BaseElement: compton.BaseElement{
			Markup:  markupSection,
			TagName: compton_atoms.Section,
		},
		r: r,
	}
}
