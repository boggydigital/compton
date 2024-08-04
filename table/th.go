package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/th.html"
	markupTh []byte
)

func NewTh() compton.Element {
	return compton.NewElement(atom.Th, markupTh)
}
