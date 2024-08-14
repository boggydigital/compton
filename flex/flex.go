package flex

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/shared"
	"io"
)

type direction string

const (
	column direction = "column"
	row    direction = "row"
)

const (
	flexElementNameTemplate = "flex-"
	rowGapAttr              = "data-row-gap"
	columnGapAttr           = "data-column-gap"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/flex-column.html"
	markupFlexColumn []byte
	//go:embed "markup/flex-row.html"
	markupFlexRow []byte
)

type Flex struct {
	compton.BaseElement
	wcr compton.Registrar
	dir direction
}

func (f *Flex) Register(w io.Writer) error {
	elementName := flexElementNameTemplate + string(f.dir)
	if f.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, f.templateFragmentWriter); err != nil {
			return err
		}
	}
	return f.BaseElement.Register(w)
}

func (f *Flex) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Direction":
		if _, err := io.WriteString(w, string(f.dir)); err != nil {
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

func NewColumn(wcr compton.Registrar) *Flex {
	return &Flex{
		BaseElement: compton.BaseElement{
			Markup:  markupFlexColumn,
			TagName: compton_atoms.ItemsCol,
		},
		wcr: wcr,
		dir: column,
	}
}

func NewRow(wcr compton.Registrar) *Flex {
	return &Flex{
		BaseElement: compton.BaseElement{
			Markup:  markupFlexRow,
			TagName: compton_atoms.ItemsCol,
		},
		wcr: wcr,
		dir: row,
	}
}
