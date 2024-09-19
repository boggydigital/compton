package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/br.html"
	markupBreak []byte
)

var Br = Break

func Break() compton.Element {
	return compton.NewElement(atom.Br, markupBreak)
}
