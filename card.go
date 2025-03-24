package compton

import (
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/issa"
	"golang.org/x/net/html/atom"
)

type CardElement struct {
	*BaseElement
	r Registrar
}

func (ce *CardElement) AppendPoster(background, placeholder, poster string, hydrated bool) *CardElement {
	if posterPlaceholder := ce.GetFirstElementByTagName(compton_atoms.Placeholder); posterPlaceholder != nil {
		if hydrated {
			hydratedPlaceholder := issa.HydrateColor(placeholder)
			posterPlaceholder.Append(IssaImageHydrated(ce.r, background, hydratedPlaceholder, poster))
		} else {
			issaImg := IssaImageDehydrated(ce.r, background, placeholder, poster)
			posterPlaceholder.Append(issaImg)
		}
	}
	return ce
}

func (ce *CardElement) AppendTitle(title string) *CardElement {
	if h3 := ce.GetFirstElementByTagName(atom.H3); h3 != nil {
		h3.Append(Text(title))
	}
	return ce
}

func (ce *CardElement) AppendProperty(title string, values ...Element) *CardElement {
	if ul := ce.GetFirstElementByTagName(atom.Ul); ul != nil {
		liProperty := Li()
		liProperty.AddClass("property")
		spanTitle := SpanText(title)
		spanTitle.AddClass("title")
		spanValues := Span()
		spanValues.AddClass("values")
		spanValues.Append(values...)
		liProperty.Append(spanTitle, spanValues)
		ul.Append(liProperty)
	}
	return ce
}

func (ce *CardElement) AppendLabels(labels ...Element) *CardElement {
	if liLables := ce.GetElementsByClassName("labels"); len(liLables) > 0 {
		liLables[0].Append(labels...)
	}
	return ce
}

func (ce *CardElement) Width(s size.Size) *CardElement {
	ce.AddClass(class.Width(s))
	return ce
}

func (ce *CardElement) WidthPixels(px float64) *CardElement {
	ce.AddClass(class.WidthPixels(px))
	return ce
}

func (ce *CardElement) Height(s size.Size) *CardElement {
	ce.AddClass(class.Height(s))
	return ce
}

func (ce *CardElement) HeightPixels(px float64) *CardElement {
	ce.AddClass(class.HeightPixels(px))
	return ce
}

func Card(r Registrar, id string) *CardElement {
	card := &CardElement{
		BaseElement: NewElement(tacMarkup(compton_atoms.Card)),
		r:           r,
	}

	card.SetAttribute("data-id", id)
	card.SetAttribute("tabindex", "-1")

	// issa-image poster slot
	card.Append(Placeholder())

	ul := Ul()
	liTitle := Li()
	liTitle.Append(H3())
	liLabels := Li()
	liLabels.AddClass("labels")
	ul.Append(liTitle, liLabels)
	card.Append(ul)

	r.RegisterStyles(DefaultStyle, compton_atoms.StyleName(compton_atoms.Card))

	return card
}
