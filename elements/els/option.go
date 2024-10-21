package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/option.html"
	markupOption []byte
)

func Option(value, label string) compton.Element {
	option := compton.NewElement(atom.Option, markupOption)
	option.SetAttribute(attr.Value, value)
	if label != "" {
		option.SetAttribute(attr.Label, label)
	}
	return option
}
