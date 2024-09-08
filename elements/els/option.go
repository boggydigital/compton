package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/option.html"
	markupOption []byte
)

func Option(value, label string) compton.Element {
	option := compton.NewElement(atom.Option, markupOption)
	option.SetAttribute(compton.ValueAttr, value)
	if label != "" {
		option.SetAttribute(compton.LabelAttr, label)
	}
	return option
}
