package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/summary.html"
	markupSummary []byte
)

func NewSummary() compton.Element {
	return compton.NewElement(atom.Summary, markupSummary)
}
