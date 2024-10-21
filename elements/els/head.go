package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/head.html"
	markupHead []byte
)

func Head() compton.Element {
	return compton.NewElement(atom.Head, markupHead)
}
