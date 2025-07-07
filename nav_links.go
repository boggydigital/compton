package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

const sectionLinksId = "section-links"

type NavTarget struct {
	Href     string
	Title    string
	Symbol   Symbol
	Selected bool
}

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

func (nle *NavLinksElement) AppendLink(r Registrar, target *NavTarget) Element {
	navListItem := ListItem()

	navLink := A(target.Href)

	if target.Selected {
		navLink.AddClass("selected")
	}

	if target.Symbol != NoSymbol {
		navLink.Append(SvgUse(r, target.Symbol))
		if target.Title != "" && target.Selected {
			navLink.Append(Text(target.Title))
		}
	} else if target.Title != "" {
		navLink.Append(Text(target.Title))
	}

	navListItem.Append(navLink)

	nle.Append(navListItem)

	return navListItem
}

func SectionsLinks(r Registrar, sections []string, sectionTitles map[string]string) Element {

	sectionNavLinks := NavLinks(r)
	sectionNavLinks.SetId(sectionLinksId)

	for _, s := range sections {
		var title string
		if t, ok := sectionTitles[s]; ok {
			title = t
		} else {
			title = s
		}
		sectionNavLinks.AppendLink(r, &NavTarget{
			Href:  "#" + title,
			Title: title,
		})
	}

	return sectionNavLinks
}
