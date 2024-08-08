package elements

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/a.html"
	markupAnchor []byte
)

func NewA(href string) compton.Element {
	anchor := compton.NewElement(atom.A, markupAnchor)
	anchor.SetAttr(compton.HrefAttr, href)
	return anchor
}

func NewAText(txt, href string) compton.Element {
	anchor := NewA(href)
	anchor.Append(NewText(txt))
	return anchor
}
