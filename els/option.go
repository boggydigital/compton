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

func NewOption(value, label string) compton.Element {
	option := compton.NewElement(atom.Option, markupOption)
	option.SetAttr(compton.ValueAttr, value)
	if label != "" {
		option.SetAttr(compton.LabelAttr, label)
	}
	return option
}
