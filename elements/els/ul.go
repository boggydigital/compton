package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/ul.html"
	markupUnorderedList []byte
)

var Ul = UnorderedList

func UnorderedList() compton.Element {
	return compton.NewElement(atom.Ul, markupUnorderedList)
}
