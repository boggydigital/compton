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
	id string
	r  Registrar
}

func (ce *CardElement) AppendPoster(background, placeholder, poster string, hydrated bool) *IssaImageElement {
	if posterPlaceholder := ce.GetFirstElementByTagName(compton_atoms.Placeholder); posterPlaceholder != nil {
		var issaImg *IssaImageElement
		if hydrated {
			hydratedPlaceholder := issa.HydrateColor(placeholder)
			issaImg = IssaImageHydrated(ce.r, background, hydratedPlaceholder, poster)
		} else {
			issaImg = IssaImageDehydrated(ce.r, background, placeholder, poster)
		}
		if issaImg != nil {
			//issaImg.SetAttribute("style", "view-transition-name:product-image-"+ce.id)
			posterPlaceholder.Append(issaImg)
		}
		return issaImg
	}
	return nil
}

func (ce *CardElement) AppendTitle(title string) *CardElement {
	if h3 := ce.GetFirstElementByTagName(atom.H3); h3 != nil {
		h3.Append(Text(title))
	}
	return ce
}

func (ce *CardElement) AppendProperty(title string, values ...Element) Element {
	if ul := ce.GetFirstElementByTagName(atom.Ul); ul != nil {
		liProperty := Li()
		liProperty.AddClass("property")
		spanTitle := SpanText(title)
		spanTitle.AddClass("title")
		spanValues := Span()
		spanValues.AddClass("values")
		spanValues.Append(values...)
		liProperty.Append(spanTitle, spanValues)
		liProperty.SetAttribute("style", "width:fit-content")
		ul.Append(liProperty)
		return liProperty
	}
	return nil
}

func (ce *CardElement) AppendBadges(badges ...Element) *CardElement {
	if liBadges := ce.GetElementsByClassName("badges"); len(liBadges) > 0 {
		liBadges[0].Append(badges...)
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
		id:          id,
		r:           r,
	}

	card.SetAttribute("data-id", id)
	card.SetAttribute("tabindex", "-1")

	// issa-image poster slot
	card.Append(Placeholder())

	ul := Ul()

	liTitle := Li()
	productTitle := H3()
	//productTitle.SetAttribute("style", "view-transition-name:product-title-"+id)
	liTitle.Append(productTitle)

	liBadges := Li()
	liBadges.AddClass("badges")
	//liBadges.SetAttribute("style", "view-transition-name:product-badges-"+id+";width:fit-content")
	ul.Append(liTitle, liBadges)

	card.Append(ul)

	r.RegisterStyles(DefaultStyle, compton_atoms.StyleName(compton_atoms.Card))

	return card
}
