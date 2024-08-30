package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/loading"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/iframe.html"
	MarkupIframe []byte
)

func NewIframe(src string) compton.Element {
	iframe := compton.NewElement(atom.Iframe, MarkupIframe)
	iframe.SetAttr(compton.SrcAttr, src)
	return iframe
}

func NewIframeLazy(src string) compton.Element {
	iframe := NewIframe(src)
	iframe.SetAttr(compton.LoadingAttr, loading.Lazy.String())
	return iframe
}
