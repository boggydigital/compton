package grid_items

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "grid-items"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/grid-items.html"
	markupGridItems []byte
	//go:embed "style/grid-items.css"
	styleGridItems []byte
)

type GridItemsElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (gie *GridItemsElement) WriteStyles(w io.Writer) error {
	if gie.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleGridItems).WriteContent(w); err != nil {
			return err
		}
	}
	return gie.BaseElement.WriteStyles(w)
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

func GridItems(wcr compton.Registrar) *GridItemsElement {
	return &GridItemsElement{
		BaseElement: compton.BaseElement{
			Markup:  markupGridItems,
			TagName: compton_atoms.GridItems,
		},
		r: wcr,
	}
}
