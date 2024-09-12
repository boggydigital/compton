package title_values

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/flex_items"
	"golang.org/x/exp/maps"
	"io"
	"slices"
)

const (
	registrationName      = "title-values"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/title-values.html"
	markupTitleValues []byte
	//go:embed "style/title-values.css"
	styleTitleValues []byte
)

type TitleValuesElement struct {
	compton.BaseElement
	r     compton.Registrar
	title compton.Element
}

func (tve *TitleValuesElement) WriteStyles(w io.Writer) error {
	if tve.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleTitleValues, styleRegistrationName).WriteContent(w); err != nil {
			return err
		}
	}
	return tve.BaseElement.WriteStyles(w)
}

func (tve *TitleValuesElement) WriteContent(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupTitleValues), w, tve.elementFragmentWriter)
}

func (tve *TitleValuesElement) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if err := tve.title.WriteContent(w); err != nil {
			return err
		}
	case compton.ContentToken:
		fallthrough
	case compton.AttributesToken:
		if err := tve.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func TitleValues(r compton.Registrar, title string) *TitleValuesElement {
	return &TitleValuesElement{
		BaseElement: compton.BaseElement{
			Markup:  markupTitleValues,
			TagName: compton_atoms.TitleValues,
		},
		r:     r,
		title: els.HeadingText(title, 3),
	}
}

func TitleValuesText(r compton.Registrar, title string, values ...string) *TitleValuesElement {
	titleValues := TitleValues(r, title)
	flexItems := flex_items.FlexItems(r, direction.Row).
		JustifyContent(align.Start).
		RowGap(size.Small).
		ColumnGap(size.Normal)

	slices.Sort(values)
	for _, value := range values {
		flexItems.Append(els.DivText(value))
	}
	titleValues.Append(flexItems)
	return titleValues
}

func TitleValuesLinks(r compton.Registrar, title string, links map[string]string, order ...string) *TitleValuesElement {
	titleValues := TitleValues(r, title)
	flexItems := flex_items.FlexItems(r, direction.Row).
		JustifyContent(align.Start).
		RowGap(size.Small).
		ColumnGap(size.Normal)

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
