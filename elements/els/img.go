package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/loading"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/img.html"
	markupImage []byte
)

var (
	Img      = Image
	ImgLazy  = ImageLazy
	ImgEager = ImageEager
)

func Image(src string) compton.Element {
	image := compton.NewElement(atom.Img, markupImage)
	if src != "" {
		image.SetAttribute(attr.Src, src)
	}
	return image
}

func ImageLazy(src string) compton.Element {
	image := Image(src)
	image.SetAttribute(attr.Loading, loading.Lazy.String())
	return image
}

func ImageEager(src string) compton.Element {
	image := Image(src)
	image.SetAttribute(attr.Loading, loading.Eager.String())
	return image
}
