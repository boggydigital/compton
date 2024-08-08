package elements

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/span.html"
	markupSpan []byte
)

func NewSpan() compton.Element {
	return compton.NewElement(atom.Span, markupSpan)
}
