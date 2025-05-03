package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

type FlexItemsElement struct {
	*BaseElement
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

func (fie *FlexItemsElement) BorderRadius(s size.Size) *FlexItemsElement {
	fie.AddClass(class.BorderRadius(s))
	return fie
}

func (fie *FlexItemsElement) Width(s size.Size) *FlexItemsElement {
	fie.AddClass(class.Width(s))
	return fie
}

func (fie *FlexItemsElement) MaxWidth(s size.Size) *FlexItemsElement {
	fie.AddClass(class.MaxWidth(s))
	return fie
}

func FlexItems(r Registrar, d direction.Direction) *FlexItemsElement {
	fie := &FlexItemsElement{
		BaseElement: NewElement(tacMarkup(compton_atoms.FlexItems)),
	}
	fie.AddClass(class.FlexDirection(d))

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.FlexItems))

	return fie
}

func FICenter(r Registrar, elements ...Element) *FlexItemsElement {
	row := FlexItems(r, direction.Row).
		JustifyContent(align.Center).
		AlignItems(align.Center)
	row.Append(elements...)
	return row
}
