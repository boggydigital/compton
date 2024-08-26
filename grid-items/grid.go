package grid_items

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/alignment"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/shared"
	"github.com/boggydigital/compton/size"
	"io"
)

const (
	gridElementName    = "grid-items"
	rowGapAttr         = "data-row-gap"
	columnGapAttr      = "data-column-gap"
	alignContentAttr   = "data-align-content"
	justifyContentAttr = "data-justify-content"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/grid-items.html"
	markupGridItems []byte
)

type Grid struct {
	compton.BaseElement
	wcr compton.Registrar
}

func (g *Grid) WriteRequirements(w io.Writer) error {
	if g.wcr.RequiresRegistration(gridElementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(gridElementName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, g.templateFragmentWriter); err != nil {
			return err
		}
	}
	return g.BaseElement.WriteRequirements(w)
}

func (g *Grid) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".HostColumnGap":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostColumnGap)); err != nil {
			return err
		}
	case ".HostRowGap":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostRowGap)); err != nil {
			return err
		}
	case ".HostAlignContent":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostAlignContent)); err != nil {
			return err
		}
	case ".HostJustifyContent":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostJustifyContent)); err != nil {
			return err
		}
	}
	return nil
}

func (g *Grid) SetRowGap(amount size.Size) *Grid {
	g.SetAttr(rowGapAttr, amount.String())
	return g
}

func (g *Grid) SetColumnGap(amount size.Size) *Grid {
	g.SetAttr(columnGapAttr, amount.String())
	return g
}

func (g *Grid) SetColumnRowGap(amount size.Size) *Grid {
	g.SetColumnGap(amount)
	g.SetRowGap(amount)
	return g
}

func (g *Grid) AlignContent(p alignment.Position) *Grid {
	g.SetAttr(alignContentAttr, p.String())
	return g
}

func (g *Grid) JustifyContent(p alignment.Position) *Grid {
	g.SetAttr(justifyContentAttr, p.String())
	return g
}

func New(wcr compton.Registrar) *Grid {
	return &Grid{
		BaseElement: compton.BaseElement{
			Markup:  markupGridItems,
			TagName: compton_atoms.GridItems,
		},
		wcr: wcr,
	}
}
