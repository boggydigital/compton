package flex_items

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/custom_elements"
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

type FlexElement struct {
	compton.BaseElement
	wcr compton.Registrar
	dir direction.Direction
}

func (f *FlexElement) WriteRequirements(w io.Writer) error {
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

func (f *FlexElement) templateFragmentWriter(t string, w io.Writer) error {
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

func (f *FlexElement) SetRowGap(amount size.Size) *FlexElement {
	f.SetAttribute(rowGapAttr, amount.String())
	return f
}

func (f *FlexElement) SetColumnGap(amount size.Size) *FlexElement {
	f.SetAttribute(columnGapAttr, amount.String())
	return f
}

func (f *FlexElement) SetGap(amount size.Size) *FlexElement {
	f.SetColumnRowGap(amount)
	f.SetRowGap(amount)
	return f
}

func (f *FlexElement) SetColumnRowGap(amount size.Size) *FlexElement {
	f.SetColumnGap(amount)
	f.SetRowGap(amount)
	return f
}

func (f *FlexElement) AlignContent(a align.Align) *FlexElement {
	f.SetAttribute(alignContentAttr, a.String())
	return f
}

func (f *FlexElement) JustifyContent(a align.Align) *FlexElement {
	f.SetAttribute(justifyContentAttr, a.String())
	return f
}

func FlexItems(wcr compton.Registrar, dir direction.Direction) *FlexElement {
	f := &FlexElement{
		BaseElement: compton.BaseElement{
			Markup:  markupFlexItems,
			TagName: compton_atoms.FlexItems,
		},
		wcr: wcr,
	}

	f.SetAttribute(flexDirectionAttr, dir.String())
	return f
}

func FlexItemsRow(wcr compton.Registrar) *FlexElement {
	return FlexItems(wcr, direction.Row)
}

func FlexItemsColumn(wcr compton.Registrar) *FlexElement {
	return FlexItems(wcr, direction.Column)
}
