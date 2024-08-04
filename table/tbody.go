package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/tbody.html"
	markupTbody []byte
)

func NewBody() compton.Element {
	return compton.NewElement(atom.Tbody, markupTbody)
}
