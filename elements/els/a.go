package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/a.html"
	markupAnchor []byte
)

func A(href string) compton.Element {
	anchor := compton.NewElement(atom.A, markupAnchor)
	anchor.SetAttribute(attr.Href, href)
	return anchor
}

func AText(txt, href string) compton.Element {
	anchor := A(href)
	anchor.Append(Text(txt))
	return anchor
}
