package nav_links

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/svg_inline"
	"io"
)

const (
	navLinksElementName = "nav-links"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/nav-links.html"
	markupNavLinks []byte
)

type NavLinksElement struct {
	compton.BaseElement
	wcr compton.Registrar
}

func (nl *NavLinksElement) WriteRequirements(w io.Writer) error {
	if nl.wcr.RequiresRegistration(navLinksElementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(navLinksElementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return nl.BaseElement.WriteRequirements(w)
}

func NavLinks(wcr compton.Registrar) *NavLinksElement {
	return &NavLinksElement{
		BaseElement: compton.BaseElement{
			Markup:  markupNavLinks,
			TagName: compton_atoms.NavLinks,
		},
		wcr: wcr,
	}
}

func NavLinksTargets(wcr compton.Registrar, targets ...*Target) *NavLinksElement {
	nl := NavLinks(wcr)

	for _, t := range targets {
		appendTarget(nl, t)

	}

	return nl
}

func appendTarget(nl *NavLinksElement, t *Target) {
	link := els.A(t.Href)

	if t.Icon != svg_inline.None {
		icon := svg_inline.SvgInline(t.Icon)
		icon.SetClass("icon")
		icon.SetAttr("title", t.Title)
		link.Append(icon)
		if t.Current {
			link.Append(els.SpanText(t.Title))
		}
	} else {
		link.Append(els.Text(t.Title))
	}
	if t.Current {
		link.SetClass("selected")
	}
	nl.Append(link)
}
