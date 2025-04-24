package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
)

type GridItemsElement struct {
	*BaseElement
}

func (gie *GridItemsElement) RowGap(sz size.Size) *GridItemsElement {
	gie.AddClass(class.RowGap(sz))
	return gie
}

func (gie *GridItemsElement) ColumnGap(sz size.Size) *GridItemsElement {
	gie.AddClass(class.ColumnGap(sz))
	return gie
}

func (gie *GridItemsElement) Gap(sz size.Size) *GridItemsElement {
	gie.ColumnGap(sz)
	gie.RowGap(sz)
	return gie
}

func (gie *GridItemsElement) AlignContent(a align.Align) *GridItemsElement {
	gie.AddClass(class.AlignContent(a))
	return gie
}

func (gie *GridItemsElement) AlignItems(a align.Align) *GridItemsElement {
	gie.AddClass(class.AlignItems(a))
	return gie
}

func (gie *GridItemsElement) JustifyContent(a align.Align) *GridItemsElement {
	gie.AddClass(class.JustifyContent(a))
	return gie
}

func (gie *GridItemsElement) JustifyItems(a align.Align) *GridItemsElement {
	gie.AddClass(class.JustifyItems(a))
	return gie
}

func GridItems(r Registrar) *GridItemsElement {
	grid := &GridItemsElement{
		BaseElement: NewElement(tacMarkup(compton_atoms.GridItems)),
	}

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.GridItems))

	return grid
}
