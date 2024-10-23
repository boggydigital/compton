package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

type FspanElement struct {
	*BaseElement
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

func (fse *FspanElement) FontWeight(w font_weight.Weight) *FspanElement {
	fse.AddClass(class.FontWeight(w))
	return fse
}

func Fspan(r Registrar, t string) *FspanElement {
	fse := &FspanElement{
		BaseElement: NewElement(tacMarkup(compton_atoms.Fspan)),
	}
	fse.Append(Text(t))

	r.RegisterStyles(comptonAtomStyle,
		compton_atoms.StyleName(compton_atoms.Fspan))

	return fse
}
