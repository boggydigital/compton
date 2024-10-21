package title_values

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
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

func (tve *TitleValuesElement) AppendValues(elements ...compton.Element) *TitleValuesElement {
	if flexItemsElements := tve.GetElementsByTagName(compton_atoms.FlexItems); len(flexItemsElements) > 0 {
		flexItemsElements[0].Append(elements...)
	} else {
		flexItems := flex_items.FlexItems(tve.r, direction.Row).
			JustifyContent(align.Start).
			RowGap(size.Small).
			ColumnGap(size.Normal)
		flexItems.Append(elements...)
		tve.Append(flexItems)
	}
	return tve
}

func (tve *TitleValuesElement) AppendTextValues(values ...string) *TitleValuesElement {
	slices.Sort(values)
	for _, value := range values {
		tve.AppendValues(els.DivText(value))
	}
	return tve
}

func (tve *TitleValuesElement) AppendLinkValues(links map[string]string, order ...string) *TitleValuesElement {
	if len(order) == 0 {
		order = maps.Keys(links)
		slices.Sort(order)
	}

	for _, key := range order {
		if links[key] != "" {
			tve.AppendValues(els.AText(key, links[key]))
		} else {
			// fallback to text if the link is empty
			tve.AppendTextValues(key)
		}
	}
	return tve
}

//func (tve *TitleValuesElement) WriteStyles(w io.Writer) error {
//	if tve.r.RequiresRegistration(styleRegistrationName) {
//		if err := els.Style(styleTitleValues, styleRegistrationName).Write(w); err != nil {
//			return err
//		}
//	}
//	return tve.BaseElement.WriteStyles(w)
//}

func (tve *TitleValuesElement) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupTitleValues), w, tve.elementFragmentWriter)
}

func (tve *TitleValuesElement) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if err := tve.title.Write(w); err != nil {
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

func (tve *TitleValuesElement) RowGap(s size.Size) *TitleValuesElement {
	tve.AddClass(class.RowGap(s))
	return tve
}

func (tve *TitleValuesElement) ForegroundColor(c color.Color) *TitleValuesElement {
	tve.AddClass(class.ForegroundColor(c))
	return tve
}

func (tve *TitleValuesElement) TitleForegroundColor(c color.Color) *TitleValuesElement {
	tve.title.AddClass(class.ForegroundColor(c))
	return tve
}

func TitleValues(r compton.Registrar, title string) *TitleValuesElement {
	tve := &TitleValuesElement{
		BaseElement: compton.BaseElement{
			Markup:  markupTitleValues,
			TagName: compton_atoms.TitleValues,
		},
		r:     r,
		title: els.HeadingText(title, 3),
	}
	tve.RowGap(size.Small)

	r.RegisterStyle(styleRegistrationName, styleTitleValues)

	return tve
}

//func TitleValuesText(r compton.Registrar, title string, values ...string) *TitleValuesElement {
//	titleValues := TitleValues(r, title)
//	flexItems := flex_items.FlexItems(r, direction.Row).
//		JustifyContent(align.Start).
//		RowGap(size.Small).
//		ColumnGap(size.Normal)
//
//	slices.Sort(values)
//	for _, value := range values {
//		flexItems.Append(els.DivText(value))
//	}
//	titleValues.Append(flexItems)
//	return titleValues
//}

//func TitleValuesLinks(r compton.Registrar, title string, links map[string]string, order ...string) *TitleValuesElement {
//	titleValues := TitleValues(r, title)
//	flexItems := flex_items.FlexItems(r, direction.Row).
//		JustifyContent(align.Start).
//		RowGap(size.Small).
//		ColumnGap(size.Normal)
//
//	if len(order) == 0 {
//		order = maps.Keys(links)
//		slices.Sort(order)
//	}
//
//	for _, key := range order {
//		flexItems.Append(els.AText(key, links[key]))
//	}
//	titleValues.Append(flexItems)
//	return titleValues
//}
