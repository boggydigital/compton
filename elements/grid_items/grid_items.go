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

	rowGapAttr         = "data-row-gap"
	columnGapAttr      = "data-column-gap"
	alignContentAttr   = "data-align-content"
	justifyContentAttr = "data-justify-content"
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

func (gi *GridItemsElement) WriteStyles(w io.Writer) error {
	if gi.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleGridItems).WriteContent(w); err != nil {
			return err
		}
	}
	return gi.BaseElement.WriteStyles(w)
}

func (g *GridItemsElement) RowGap(sz size.Size) *GridItemsElement {
	g.AddClass(class.RowGap(sz))
	return g
}

func (g *GridItemsElement) ColumnGap(sz size.Size) *GridItemsElement {
	g.AddClass(class.ColumnGap(sz))
	return g
}

func (g *GridItemsElement) Gap(sz size.Size) *GridItemsElement {
	g.ColumnGap(sz)
	g.RowGap(sz)
	return g
}

func (g *GridItemsElement) AlignContent(a align.Align) *GridItemsElement {
	g.AddClass(class.AlignContent(a))
	return g
}

func (g *GridItemsElement) AlignItems(a align.Align) *GridItemsElement {
	g.AddClass(class.AlignItems(a))
	return g
}

func (g *GridItemsElement) JustifyContent(a align.Align) *GridItemsElement {
	g.AddClass(class.JustifyContent(a))
	return g
}

func (g *GridItemsElement) JustifyItems(a align.Align) *GridItemsElement {
	g.AddClass(class.JustifyItems(a))
	return g
}

func GridItems(wcr compton.Registrar) *GridItemsElement {
	gi := &GridItemsElement{
		BaseElement: compton.BaseElement{
			Markup:  markupGridItems,
			TagName: compton_atoms.GridItems,
		},
		r: wcr,
	}
	return gi
}
