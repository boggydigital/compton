package title_values

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
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

type TitleValues struct {
	compton.BaseElement
	wcr   compton.Registrar
	title compton.Element
}

func (tv *TitleValues) WriteRequirements(w io.Writer) error {
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

func (tv *TitleValues) WriteContent(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupTitleValues), w, tv.elementFragmentWriter)
}

func (tv *TitleValues) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if err := tv.title.WriteContent(w); err != nil {
			return err
		}
		//if _, err := io.WriteString(w, tv.title); err != nil {
		//	return err
		//}
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

func New(wcr compton.Registrar, title string) *TitleValues {
	return &TitleValues{
		BaseElement: compton.BaseElement{
			Markup:  markupTitleValues,
			TagName: compton_atoms.TitleValues,
		},
		wcr:   wcr,
		title: els.NewHeadingText(title, 3),
	}
}

func NewText(wcr compton.Registrar, title string, values ...string) *TitleValues {
	titleValues := New(wcr, title)
	flexItems := flex_items.New(wcr, direction.Row).
		SetRowGap(size.Normal).
		SetColumnGap(size.Normal)

	slices.Sort(values)
	for _, value := range values {
		flexItems.Append(els.NewDivText(value))
	}
	titleValues.Append(flexItems)
	return titleValues
}

func NewLinks(wcr compton.Registrar, title string, links map[string]string, order ...string) *TitleValues {
	titleValues := New(wcr, title)
	flexItems := flex_items.New(wcr, direction.Row).
		SetRowGap(size.Normal).
		SetColumnGap(size.Normal)

	if len(order) == 0 {
		order = maps.Keys(links)
		slices.Sort(order)
	}

	for _, key := range order {
		flexItems.Append(els.NewAText(key, links[key]))
	}
	titleValues.Append(flexItems)
	return titleValues
}
