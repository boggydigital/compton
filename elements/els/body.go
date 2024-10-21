package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/body.html"
	markupBody []byte
)

func Body() compton.Element {
	return compton.NewElement(atom.Body, markupBody)
}
