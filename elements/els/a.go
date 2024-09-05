package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/a.html"
	markupAnchor []byte
)

func A(href string) compton.Element {
	anchor := compton.NewElement(atom.A, markupAnchor)
	anchor.SetAttr(compton.HrefAttr, href)
	return anchor
}

func AText(txt, href string) compton.Element {
	anchor := A(href)
	anchor.Append(Text(txt))
	return anchor
}
