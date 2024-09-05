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

func Span() compton.Element {
	return compton.NewElement(atom.Span, markupSpan)
}

func SpanText(txt string) compton.Element {
	span := Span()
	span.Append(Text(txt))
	return span
}
