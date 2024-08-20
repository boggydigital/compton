package flex_items

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/anchors"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/directions"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/shared"
	"io"
)

const (
	flexElementName    = "flex-items"
	rowGapAttr         = "data-row-gap"
	columnGapAttr      = "data-column-gap"
	alignContentAttr   = "data-align-content"
	justifyContentAttr = "data-justify-content"
	flexDirectionAttr  = "data-flex-direction"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/flex-items.html"
	markupFlexItems []byte
)

type Flex struct {
	compton.BaseElement
	wcr compton.Registrar
	dir directions.Direction
}

func (f *Flex) WriteRequirements(w io.Writer) error {
	if f.wcr.RequiresRegistration(flexElementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(flexElementName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, f.templateFragmentWriter); err != nil {
			return err
		}
	}
	return f.BaseElement.WriteRequirements(w)
}

func (f *Flex) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Direction":
		if _, err := io.WriteString(w, f.dir.String()); err != nil {
			return err
		}
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
	case ".HostFlexDirection":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostFlexDirection)); err != nil {
			return err
		}
	}
	return nil
}

func (f *Flex) SetRowGap(amount measures.Unit) *Flex {
	f.SetAttr(rowGapAttr, amount.String())
	return f
}

func (f *Flex) SetColumnGap(amount measures.Unit) *Flex {
	f.SetAttr(columnGapAttr, amount.String())
	return f
}

func (f *Flex) SetColumnRowGap(amount measures.Unit) *Flex {
	f.SetColumnGap(amount)
	f.SetRowGap(amount)
	return f
}

func (f *Flex) AlignContent(p anchors.Position) *Flex {
	f.SetAttr(alignContentAttr, p.String())
	return f
}

func (f *Flex) JustifyContent(p anchors.Position) *Flex {
	f.SetAttr(justifyContentAttr, p.String())
	return f
}

func New(wcr compton.Registrar, dir directions.Direction) *Flex {
	f := &Flex{
		BaseElement: compton.BaseElement{
			Markup:  markupFlexItems,
			TagName: compton_atoms.FlexItems,
		},
		wcr: wcr,
	}

	f.SetAttr(flexDirectionAttr, dir.String())
	return f
}
