package nav_links

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/svg_inline"
	"io"
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
	r compton.Registrar
}

func (nle *NavLinksElement) WriteStyles(w io.Writer) error {
	if nle.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleNavLinks).WriteContent(w); err != nil {
			return err
		}
	}
	return nle.BaseElement.WriteStyles(w)
}

func NavLinks(r compton.Registrar) *NavLinksElement {
	return &NavLinksElement{
		BaseElement: compton.BaseElement{
			Markup:  markupNavLinks,
			TagName: compton_atoms.NavLinks,
		},
		r: r,
	}
}

func NavLinksTargets(r compton.Registrar, targets ...*Target) *NavLinksElement {
	nl := NavLinks(r)
	for _, t := range targets {
		appendTarget(nl, t)
	}
	return nl
}

func appendTarget(nl *NavLinksElement, t *Target) {
	link := els.A(t.Href)

	if t.Icon != svg_inline.None {
		icon := svg_inline.SvgInline(t.Icon)
		icon.AddClass("icon")
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
	nl.Append(link)
}
