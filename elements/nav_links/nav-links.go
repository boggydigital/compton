package nav_links

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/svg_use"
)

const (
	registrationName      = "nav-links"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/nav-links.html"
	markupNavLinks []byte
	//go:embed "style/nav-links.css"
	styleNavLinks []byte
)

type NavLinksElement struct {
	compton.BaseElement
	//r compton.Registrar
}

//func (nle *NavLinksElement) WriteStyles(w io.Writer) error {
//	if nle.r.RequiresRegistration(styleRegistrationName) {
//		if err := els.Style(styleNavLinks, styleRegistrationName).Write(w); err != nil {
//			return err
//		}
//	}
//	return nle.BaseElement.WriteStyles(w)
//}

func NavLinks(r compton.Registrar) *NavLinksElement {
	navLinks := &NavLinksElement{
		BaseElement: compton.BaseElement{
			Markup:  markupNavLinks,
			TagName: compton_atoms.NavLinks,
		},
		//r: r,
	}

	r.RegisterStyle(styleRegistrationName, styleNavLinks)

	return navLinks

}

func NavLinksTargets(r compton.Registrar, targets ...*Target) *NavLinksElement {
	nl := NavLinks(r)
	for _, t := range targets {
		appendTarget(r, nl, t)
	}
	return nl
}

func appendTarget(r compton.Registrar, nl *NavLinksElement, t *Target) {
	li := els.ListItem()
	link := els.A(t.Href)

	if t.Icon != svg_use.None {
		icon := svg_use.SvgUse(r, t.Icon)
		//icon.AddClass("icon")
		icon.SetAttribute("title", t.Title)
		link.Append(icon)
		if t.Current {
			link.Append(els.SpanText(t.Title))
		}
	} else {
		link.Append(els.Text(t.Title))
	}
	if t.Current {
		link.AddClass("selected")
	}
	li.Append(link)
	nl.Append(li)
}
