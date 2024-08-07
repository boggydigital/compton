package anchor

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/text"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/a.html"
	markupAnchor []byte
)

func New(href string) compton.Element {
	anchor := compton.NewElement(atom.A, markupAnchor)
	anchor.SetAttr(compton.HrefAttr, href)
	return anchor
}

func NewText(txt, href string) compton.Element {
	anchor := New(href)
	anchor.Append(text.New(txt))
	return anchor
}
