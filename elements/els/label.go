package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/label.html"
	markupLabel []byte
)

func NewLabel(forInput string) compton.Element {
	label := compton.NewElement(atom.Label, markupLabel)
	label.SetAttr(compton.ForAttr, forInput)
	return label
}
