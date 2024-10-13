package flex_items

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "flex-items"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/flex-items.html"
	markupFlexItems []byte
	//go:embed "style/flex-items.css"
	styleFlexItems []byte
)

type FlexItemsElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (fie *FlexItemsElement) WriteStyles(w io.Writer) error {
	if fie.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleFlexItems, styleRegistrationName).WriteContent(w); err != nil {
			return err
		}
	}
	return fie.BaseElement.WriteStyles(w)
}

func (fie *FlexItemsElement) RowGap(sz size.Size) *FlexItemsElement {
	fie.AddClass(class.RowGap(sz))
	return fie
}

func (fie *FlexItemsElement) ColumnGap(sz size.Size) *FlexItemsElement {
	fie.AddClass(class.ColumnGap(sz))
	return fie
}

func (fie *FlexItemsElement) Gap(sz size.Size) *FlexItemsElement {
	fie.ColumnGap(sz)
	fie.RowGap(sz)
	return fie
}

func (fie *FlexItemsElement) AlignContent(a align.Align) *FlexItemsElement {
	fie.AddClass(class.AlignContent(a))
	return fie
}

func (fie *FlexItemsElement) AlignItems(a align.Align) *FlexItemsElement {
	fie.AddClass(class.AlignItems(a))
	return fie
}

func (fie *FlexItemsElement) JustifyContent(a align.Align) *FlexItemsElement {
	fie.AddClass(class.JustifyContent(a))
	return fie
}

func (fie *FlexItemsElement) JustifyItems(a align.Align) *FlexItemsElement {
	fie.AddClass(class.JustifyItems(a))
	return fie
}

func (fie *FlexItemsElement) FontSize(s size.Size) *FlexItemsElement {
	fie.AddClass(class.FontSize(s))
	return fie
}

func (fie *FlexItemsElement) FontWeight(w font_weight.Weight) *FlexItemsElement {
	fie.AddClass(class.FontWeight(w))
	return fie
}

func (fie *FlexItemsElement) ForegroundColor(c color.Color) *FlexItemsElement {
	fie.AddClass(class.ForegroundColor(c))
	return fie
}

func (fie *FlexItemsElement) BackgroundColor(c color.Color) *FlexItemsElement {
	fie.AddClass(class.BackgroundColor(c))
	return fie
}

func FlexItems(r compton.Registrar, d direction.Direction) *FlexItemsElement {
	fie := &FlexItemsElement{
		BaseElement: compton.BaseElement{
			Markup:  markupFlexItems,
			TagName: compton_atoms.FlexItems,
		},
		r: r,
	}
	fie.AddClass(class.FlexDirection(d))
	return fie
}

func Center(r compton.Registrar, elements ...compton.Element) *FlexItemsElement {
	row := FlexItems(r, direction.Row).
		JustifyContent(align.Center).
		AlignItems(align.Center)
	row.Append(elements...)
	return row
}
