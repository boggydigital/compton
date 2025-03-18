package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

const sectionLinksId = "section-links"

const SectionLinksTitle = "&#x2935;" // ARROW POINTING RIGHTWARDS THEN CURVING DOWNWARDS

type NavLinksElement struct {
	*BaseElement
}

func NavLinks(r Registrar) *NavLinksElement {
	navLinks := &NavLinksElement{
		BaseElement: NewElement(atomsEmbedMarkup(compton_atoms.NavLinks, DefaultMarkup)),
	}

	r.RegisterStyles(DefaultStyle,
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

func SectionsLinks(r Registrar, sections []string, sectionTitles map[string]string) Element {

	sectionLinks := make(map[string]string)
	sectionsOrder := make([]string, 0, len(sections))
	for _, s := range sections {
		var title string
		if t, ok := sectionTitles[s]; ok {
			title = t
		} else {
			title = s
		}
		sectionLinks[title] = "#" + title
		sectionsOrder = append(sectionsOrder, title)
	}

	targets := TextLinks(sectionLinks, "", sectionsOrder...)

	psl := NavLinksTargets(r, targets...)
	psl.SetId(sectionLinksId)

	return psl

}
