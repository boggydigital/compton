package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

type NavLinksElement struct {
	*BaseElement
}

func NavLinks(r Registrar) *NavLinksElement {
	navLinks := &NavLinksElement{
		BaseElement: NewElement(atomsEmbedMarkup(compton_atoms.NavLinks, comptonAtomsMarkup)),
	}

	r.RegisterStyles(comptonAtomStyle,
		compton_atoms.StyleName(compton_atoms.NavLinks))

	return navLinks
}

func NavLinksTargets(r Registrar, targets ...*Target) *NavLinksElement {
	nl := NavLinks(r)
	for _, t := range targets {
		appendTarget(r, nl, t)
	}
	return nl
}

func appendTarget(r Registrar, nl *NavLinksElement, t *Target) {
	li := ListItem()
	link := A(t.Href)

	if t.Icon != None {
		icon := SvgUse(r, t.Icon)
		icon.SetAttribute("title", t.Title)
		link.Append(icon)
		if t.Current {
			icon.ForegroundColor(color.Background)
			link.Append(SpanText(t.Title))
		}
	} else {
		link.Append(Text(t.Title))
	}
	if t.Current {
		link.AddClass("selected")
	}
	li.Append(link)
	nl.Append(li)
}
