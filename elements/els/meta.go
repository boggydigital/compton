package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/meta.html"
	markupMeta []byte
)

func Meta(kv map[string]string) compton.Element {
	meta := compton.NewElement(atom.Meta, markupMeta)
	for k, v := range kv {
		meta.SetAttribute(k, v)
	}
	return meta
}
