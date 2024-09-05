package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/ol.html"
	markupOrderedList []byte
)

var Ol = OrderedList

func OrderedList() compton.Element {
	return compton.NewElement(atom.Ol, markupOrderedList)
}
