package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/link.html"
	markupLink []byte
)

func Link(kv map[string]string) compton.Element {
	link := compton.NewElement(atom.Link, markupLink)
	for k, v := range kv {
		link.SetAttribute(k, v)
	}
	return link
}
