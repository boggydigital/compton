package els

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

func NewSpanText(txt string) compton.Element {
	span := NewSpan()
	span.Append(NewText(txt))
	return span
}
