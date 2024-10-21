package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/loading"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/iframe.html"
	MarkupIframe []byte
)

func Iframe(src string) compton.Element {
	iframe := compton.NewElement(atom.Iframe, MarkupIframe)
	iframe.SetAttribute(attr.Src, src)
	return iframe
}

func IframeLazy(src string) compton.Element {
	iframe := Iframe(src)
	iframe.SetAttribute(attr.Loading, loading.Lazy.String())
	return iframe
}
