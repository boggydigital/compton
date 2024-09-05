package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/tfoot.html"
	markupTfoot []byte
)

func Tfoot() compton.Element {
	return compton.NewElement(atom.Tfoot, markupTfoot)
}
