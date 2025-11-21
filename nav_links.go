package compton

import (
	_ "embed"

	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"golang.org/x/net/html/atom"
)

const sectionLinksId = "section-links"

//go:embed "script/nav_link_submit_form.js"
var navLinkSubmitForm []byte

type NavTarget struct {
	Href  string
	Title string
	//Symbol   Symbol
	IconElement Element
	Selected    bool
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

	//if target.Symbol != NoSymbol {
	//	navLink.Append(SvgUse(r, target.Symbol))
	//	if target.Title != "" && target.Selected {
	//		navLink.Append(Text(target.Title))
	//	}
	//}
	if target.IconElement != nil {
		navLink.Append(target.IconElement)
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

func (nle *NavLinksElement) AppendSubmitLink(r Registrar, target *NavTarget) Element {

	r.RegisterDeferrals("nav-link-submit-script", ScriptAsync(navLinkSubmitForm))

	submitListItem := nle.AppendLink(r, target)
	if link := submitListItem.GetFirstElementByTagName(atom.A); link != nil {
		link.AddClass("submit")
	}

	return submitListItem
}

func (nle *NavLinksElement) Width(s size.Size) *NavLinksElement {
	nle.AddClass(class.Width(s))
	return nle
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
