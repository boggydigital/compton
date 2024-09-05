package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/time.html"
	markupTime []byte
)

func Time() compton.Element {
	return compton.NewElement(atom.Time, markupTime)
}

func TimeText(txt string) compton.Element {
	tm := Time()
	tm.Append(Text(txt))
	return tm
}
