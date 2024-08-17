package nav_links

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/els"
	"golang.org/x/exp/maps"
	"io"
	"sort"
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

type NavLinks struct {
	compton.BaseElement
	wcr compton.Registrar
}

func (nl *NavLinks) Register(w io.Writer) error {
	if nl.wcr.RequiresRegistration(navLinksElementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(navLinksElementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return nl.BaseElement.Register(w)
}

func New(wcr compton.Registrar) *NavLinks {
	return &NavLinks{
		BaseElement: compton.BaseElement{
			Markup:  markupNavLinks,
			TagName: compton_atoms.GridItems,
		},
		wcr: wcr,
	}
}

func NewLinks(wcr compton.Registrar, links map[string]string, order ...string) *NavLinks {
	nl := New(wcr)

	if len(order) == 0 {
		order = maps.Keys(links)
		sort.Strings(order)
	}

	for _, key := range order {
		nl.Append(els.NewAText(key, links[key]))
	}
	return nl
}
