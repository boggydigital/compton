package title_values

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/flex_items"
	"golang.org/x/exp/maps"
	"io"
	"slices"
)

const (
	elementName = "title-values"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/title-values.html"
	markupTitleValues []byte
)

type TitleValuesElement struct {
	compton.BaseElement
	wcr   compton.Registrar
	title compton.Element
}

func (tv *TitleValuesElement) WriteRequirements(w io.Writer) error {
	if tv.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return tv.BaseElement.WriteRequirements(w)
}

func (tv *TitleValuesElement) WriteContent(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupTitleValues), w, tv.elementFragmentWriter)
}

func (tv *TitleValuesElement) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if err := tv.title.WriteContent(w); err != nil {
			return err
		}
	case compton.ContentToken:
		fallthrough
	case compton.AttributesToken:
		if err := tv.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func TitleValues(wcr compton.Registrar, title string) *TitleValuesElement {
	return &TitleValuesElement{
		BaseElement: compton.BaseElement{
			Markup:  markupTitleValues,
			TagName: compton_atoms.TitleValues,
		},
		wcr:   wcr,
		title: els.HeadingText(title, 3),
	}
}

func TitleValuesText(r compton.Registrar, title string, values ...string) *TitleValuesElement {
	titleValues := TitleValues(r, title)
	flexItems := flex_items.FlexItems(r, direction.Row)
	//RowGap(size.Normal).
	//ColumnGap(size.Normal)

	slices.Sort(values)
	for _, value := range values {
		flexItems.Append(els.DivText(value))
	}
	titleValues.Append(flexItems)
	return titleValues
}

func TitleValuesLinks(r compton.Registrar, title string, links map[string]string, order ...string) *TitleValuesElement {
	titleValues := TitleValues(r, title)
	flexItems := flex_items.FlexItems(r, direction.Row)
	//SetRowGap(size.Normal).
	//SetColumnGap(size.Normal)

	if len(order) == 0 {
		order = maps.Keys(links)
		slices.Sort(order)
	}

	for _, key := range order {
		flexItems.Append(els.AText(key, links[key]))
	}
	titleValues.Append(flexItems)
	return titleValues
}
