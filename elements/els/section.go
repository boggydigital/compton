package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/section.html"
	markupSection []byte
)

func Section() compton.Element {
	return compton.NewElement(atom.Section, markupSection)
}
