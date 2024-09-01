package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/tr.html"
	markupTr []byte
)

func NewTr() compton.Element {
	return compton.NewElement(atom.Tr, markupTr)
}
