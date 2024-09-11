package fspan

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
	registrationName      = "fspan"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/fspan.html"
	markupFspan []byte
	//go:embed "style/fspan.css"
	styleFspan []byte
)

type FspanElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (fse *FspanElement) WriteStyles(w io.Writer) error {
	if fse.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleFspan).WriteContent(w); err != nil {
			return err
		}
	}
	return fse.BaseElement.WriteStyles(w)
}

func (fse *FspanElement) ForegroundColor(c color.Color) *FspanElement {
	fse.AddClass(class.ForegroundColor(c))
	return fse
}

func (fse *FspanElement) BackgroundColor(c color.Color) *FspanElement {
	fse.AddClass(class.BackgroundColor(c))
	return fse
}

func (fse *FspanElement) FontSize(s size.Size) *FspanElement {
	fse.AddClass(class.FontSize(s))
	return fse
}

func (fse *FspanElement) FontWeight(w weight.Weight) *FspanElement {
	fse.AddClass(class.FontWeight(w))
	return fse
}

func Text(r compton.Registrar, t string) *FspanElement {
	fse := &FspanElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.Fspan,
			Markup:  markupFspan,
		},
		r: r,
	}
	fse.Append(els.Text(t))
	return fse
}
