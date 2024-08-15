package title_values

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/directions"
	"github.com/boggydigital/compton/els"
	flex_items "github.com/boggydigital/compton/flex-items"
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
	title string
}

func (tv *TitleValues) Register(w io.Writer) error {
	if tv.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return tv.BaseElement.Register(w)
}

func (tv *TitleValues) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupTitleValues), w, tv.elementFragmentWriter)
}

func (tv *TitleValues) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if _, err := io.WriteString(w, tv.title); err != nil {
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

func New(wcr compton.Registrar, title string) *TitleValues {
	return &TitleValues{
		BaseElement: compton.BaseElement{
			Markup:  markupTitleValues,
			TagName: compton_atoms.DetailsOpen,
		},
		wcr:   wcr,
		title: title,
	}
}

func NewText(wcr compton.Registrar, title string, values ...string) *TitleValues {
	titleValues := New(wcr, title)
	flexItems := flex_items.New(wcr, directions.Row)
	slices.Sort(values)
	for _, value := range values {
		flexItems.Append(els.NewDivText(value))
	}
	titleValues.Append(flexItems)
	return titleValues
}

func NewLinks(wcr compton.Registrar, title string, links map[string]string) *TitleValues {
	titleValues := New(wcr, title)
	flexItems := flex_items.New(wcr, directions.Row)
	keys := maps.Keys(links)
	slices.Sort(keys)
	for _, key := range keys {
		flexItems.Append(els.NewAText(key, links[key]))
	}
	titleValues.Append(flexItems)
	return titleValues
}
