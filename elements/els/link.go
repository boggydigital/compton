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

func Link() compton.Element {
	return compton.NewElement(atom.Link, markupLink)
}
