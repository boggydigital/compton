package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/thead.html"
	markupThead []byte
)

func NewHead() compton.Element {
	return compton.NewElement(atom.Thead, markupThead)
}
