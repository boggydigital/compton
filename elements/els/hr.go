package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/hr.html"
	markupHorizontalRule []byte
)

var Hr = HorizontalRule

func HorizontalRule() compton.Element {
	return compton.NewElement(atom.Hr, markupHorizontalRule)
}
