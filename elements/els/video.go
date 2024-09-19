package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/video.html"
	markupVideo []byte
)

func Video(src string) compton.Element {
	video := compton.NewElement(atom.Video, markupVideo)
	if src != "" {
		video.SetAttribute(compton.SrcAttr, src)
	}
	return video
}
