package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/img.html"
	markupImage []byte
)

func NewImage(src string) compton.Element {
	image := compton.NewElement(atom.Img, markupAnchor)
	image.SetAttr(compton.SrcAttr, src)
	return image
}