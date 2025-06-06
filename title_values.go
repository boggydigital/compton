package compton

import (
	"bytes"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"io"
	"maps"
	"slices"
)

const LinkTargetTop = "_top"

type TitleValuesElement struct {
	*BaseElement
	r          Registrar
	title      Element
	linkTarget string
}

func (tve *TitleValuesElement) AppendValues(elements ...Element) *TitleValuesElement {
	if flexItemsElements := tve.GetElementsByTagName(compton_atoms.FlexItems); len(flexItemsElements) > 0 {
		flexItemsElements[0].Append(elements...)
	} else {
		flexItems := FlexItems(tve.r, direction.Row).
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
		tve.AppendValues(DivText(value))
	}
	return tve
}

func (tve *TitleValuesElement) SetLinksTarget(target string) *TitleValuesElement {
	tve.linkTarget = target
	return tve
}

func (tve *TitleValuesElement) AppendLinkValues(links map[string]string, order ...string) *TitleValuesElement {
	if len(order) == 0 {
		order = slices.Sorted(maps.Keys(links))
	}

	for _, key := range order {
		if links[key] != "" {
			link := AText(key, links[key])
			if tve.linkTarget != "" {
				link.SetAttribute(attr.Target, tve.linkTarget)
			}
			tve.AppendValues(link)
		} else {
			// fallback to text if the link is empty
			tve.AppendTextValues(key)
		}
	}
	return tve
}

func (tve *TitleValuesElement) Write(w io.Writer) error {
	bts, err := tve.BaseElement.MarkupProvider.GetMarkup()
	if err != nil {
		return err
	}
	return WriteContents(bytes.NewReader(bts), w, tve.elementFragmentWriter)
}

func (tve *TitleValuesElement) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if err := tve.title.Write(w); err != nil {
			return err
		}
	case ContentToken:
		fallthrough
	case AttributesToken:
		if err := tve.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return ErrUnknownToken(t)
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

func TitleValues(r Registrar, title string) *TitleValuesElement {
	tve := &TitleValuesElement{
		BaseElement: NewElement(atomsEmbedMarkup(compton_atoms.TitleValues, DefaultMarkup)),
		r:           r,
		title:       HeadingText(title, 3),
	}

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.TitleValues))

	return tve
}
